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
	"sync/atomic"
	"syscall"

	"github.com/murkland/signor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	listenAddr = flag.String("listen_addr", "[::]:12345", "address to listen on")
)

type session struct {
	id            string
	refcount      int32
	offerSDP      string
	answerSDPChan chan string
}

type server struct {
	pb.UnimplementedSessionServiceServer

	sessions   map[string]*session
	sessionsMu sync.Mutex
}

func (s *server) Negotiate(stream pb.SessionService_NegotiateServer) error {
	var sess *session

	defer func() {
		if sess == nil {
			return
		}

		s.sessionsMu.Lock()
		defer s.sessionsMu.Unlock()
		if atomic.AddInt32(&sess.refcount, -1) <= 0 {
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
					sess = &session{id: p.Start.SessionId, refcount: 0, answerSDPChan: make(chan string)}
					s.sessions[sess.id] = sess
				}
			}()
			atomic.AddInt32(&sess.refcount, 1)

			if sess.offerSDP != "" {
				stream.Send(&pb.NegotiateResponse{
					Which: &pb.NegotiateResponse_Offer_{
						Offer: &pb.NegotiateResponse_Offer{
							Sdp: sess.offerSDP,
						},
					},
				})
			} else {
				sess.offerSDP = p.Start.OfferSdp
				select {
				case answerSDP := <-sess.answerSDPChan:
					stream.Send(&pb.NegotiateResponse{
						Which: &pb.NegotiateResponse_Answer_{
							Answer: &pb.NegotiateResponse_Answer{
								Sdp: answerSDP,
							},
						},
					})
					return nil
				case <-stream.Context().Done():
					return stream.Context().Err()
				}
			}
		case *pb.NegotiateRequest_Answer_:
			if sess == nil {
				return grpc.Errorf(codes.FailedPrecondition, "did not receive start packet")
			}
			select {
			case sess.answerSDPChan <- p.Answer.Sdp:
			case <-stream.Context().Done():
				return stream.Context().Err()
			}
			return nil
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
