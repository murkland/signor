package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/murkland/signor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	listenAddr = flag.String("listen_addr", "[::]:12345", "address to listen on")
)

type session struct {
	refcount   int
	refcountMu sync.Mutex

	id       string
	offerSDP string
	streams  [2]pb.SessionService_NegotiateServer
}

type server struct {
	pb.UnimplementedSessionServiceServer

	sessions   map[string]*session
	sessionsMu sync.Mutex
}

func (s *server) Negotiate(stream pb.SessionService_NegotiateServer) error {
	var sess *session
	var me int

	defer func() {
		if sess == nil {
			return
		}

		s.sessionsMu.Lock()
		defer s.sessionsMu.Unlock()

		sess.refcountMu.Lock()
		sess.refcount--
		sess.refcountMu.Unlock()

		if sess.refcount <= 0 {
			delete(s.sessions, sess.id)
		}
	}()
	for {
		in, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		switch p := in.Which.(type) {
		case *pb.NegotiateRequest_Start_:
			func() {
				s.sessionsMu.Lock()
				defer s.sessionsMu.Unlock()
				sess = s.sessions[p.Start.SessionId]
				if sess == nil {
					sess = &session{
						id:       p.Start.SessionId,
						offerSDP: p.Start.OfferSdp,
						refcount: 0,
					}
					s.sessions[sess.id] = sess
				}
			}()

			if err := func() error {
				sess.refcountMu.Lock()
				defer sess.refcountMu.Unlock()
				if sess.refcount >= 2 {
					return grpc.Errorf(codes.FailedPrecondition, "too many clients on this session")
				}
				me = sess.refcount
				sess.refcount++
				return nil
			}(); err != nil {
				return err
			}

			sess.streams[me] = stream

			// This is the polite peer.
			if me == 1 {
				if err := stream.Send(&pb.NegotiateResponse{
					Which: &pb.NegotiateResponse_Offer_{
						Offer: &pb.NegotiateResponse_Offer{
							Sdp: sess.offerSDP,
						},
					},
				}); err != nil {
					return err
				}
			}
		case *pb.NegotiateRequest_Answer_:
			if sess == nil {
				return grpc.Errorf(codes.FailedPrecondition, "did not receive start packet")
			}

			if me != 1 {
				return grpc.Errorf(codes.FailedPrecondition, "only the impolite peer may answer")
			}

			if err := sess.streams[0].Send(&pb.NegotiateResponse{
				Which: &pb.NegotiateResponse_Answer_{
					Answer: &pb.NegotiateResponse_Answer{
						Sdp: p.Answer.Sdp,
					},
				},
			}); err != nil {
				return err
			}
		case *pb.NegotiateRequest_IceCandidate:
			if sess == nil {
				return grpc.Errorf(codes.FailedPrecondition, "did not receive start packet")
			}

			peerStream := sess.streams[1-me]
			if peerStream == nil {
				return grpc.Errorf(codes.FailedPrecondition, "no peer stream")
			}

			if err := peerStream.Send(&pb.NegotiateResponse{
				Which: &pb.NegotiateResponse_IceCandidate{
					IceCandidate: &pb.NegotiateResponse_ICECandidate{
						IceCandidate: p.IceCandidate.IceCandidate,
					},
				},
			}); err != nil {
				return err
			}
		}
	}
}

func main() {
	flag.Parse()

	s := &server{
		sessions: map[string]*session{},
	}

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("net.Listen(): %s", err)
	}
	log.Printf("listening on %s", lis.Addr())

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGUSR1)
	go func() {
		for range sigChan {
			s.sessionsMu.Lock()
			log.Printf("active sessions: %d", len(s.sessions))
			s.sessionsMu.Unlock()
		}
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterSessionServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
