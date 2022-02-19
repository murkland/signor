package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/cors"
	"github.com/twitchtv/twirp"
	"github.com/yumland/signor/pb"
)

var (
	listenAddr = flag.String("listen_addr", "[::]:12345", "address to listen on")
)

type session struct {
	offerSDP      string
	answerSDPChan chan string
}

type server struct {
	sessions   map[string]*session
	sessionsMu sync.Mutex
}

func (s *server) Offer(ctx context.Context, req *pb.OfferRequest) (*pb.OfferResponse, error) {
	sess := &session{
		offerSDP:      req.MyOfferSdp,
		answerSDPChan: make(chan string),
	}

	if err := (func() error {
		s.sessionsMu.Lock()
		defer s.sessionsMu.Unlock()
		_, ok := s.sessions[string(req.SessionId)]
		if !ok {
			return twirp.AlreadyExists.Error("session already exists")
		}

		s.sessions[string(req.SessionId)] = sess
		return nil
	})(); err != nil {
		return nil, err
	}

	defer (func() {
		s.sessionsMu.Lock()
		defer s.sessionsMu.Unlock()
		delete(s.sessions, string(req.SessionId))
	})()

	var answerSDP string
	select {
	case answerSDP = <-sess.answerSDPChan:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return &pb.OfferResponse{
		TheirAnswerSdp: answerSDP,
	}, nil
}

func (s *server) GetOffer(ctx context.Context, req *pb.GetOfferRequest) (*pb.GetOfferResponse, error) {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	sess := s.sessions[string(req.SessionId)]
	if sess == nil {
		return nil, twirp.NotFound.Error("no such session")
	}

	return &pb.GetOfferResponse{
		TheirOfferSdp: sess.offerSDP,
	}, nil
}

func (s *server) Answer(ctx context.Context, req *pb.AnswerRequest) (*pb.AnswerResponse, error) {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	sess := s.sessions[string(req.SessionId)]
	if sess == nil {
		return nil, twirp.NotFound.Error("no such session")
	}

	select {
	case sess.answerSDPChan <- req.MyAnswerSdp:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	delete(s.sessions, string(req.SessionId))
	return &pb.AnswerResponse{}, nil
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

	http.Serve(lis, cors.New(cors.Options{
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type", "Twirp-Version"},
	}).Handler(pb.NewSessionServiceServer(s)))
}
