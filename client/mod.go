package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/murkland/signor/pb"
	"github.com/pion/webrtc/v3"
	"google.golang.org/grpc"
)

type Client struct {
	rpcClient pb.SessionServiceClient
}

func New(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{pb.NewSessionServiceClient(conn)}, nil
}

type ConnectionSide int

const (
	ConnectionSideUnknown ConnectionSide = 0
	ConnectionSideOfferer                = iota
	ConnectionSideAnswerer
)

func (c *Client) Connect(ctx context.Context, sessionID string, makePeerConn func() (*webrtc.PeerConnection, error)) (*webrtc.PeerConnection, ConnectionSide, error) {
	peerConn, err := makePeerConn()
	if err != nil {
		return nil, ConnectionSideUnknown, err
	}

	offer, err := peerConn.CreateOffer(nil)
	if err != nil {
		return nil, ConnectionSideUnknown, err
	}
	if err = peerConn.SetLocalDescription(offer); err != nil {
		return nil, ConnectionSideUnknown, err
	}

	negotiation, err := c.rpcClient.Negotiate(ctx)
	if err != nil {
		return nil, ConnectionSideUnknown, err
	}

	if err := negotiation.Send(&pb.NegotiateRequest{
		Which: &pb.NegotiateRequest_Start_{
			Start: &pb.NegotiateRequest_Start{
				SessionId: sessionID,
				OfferSdp:  peerConn.LocalDescription().SDP,
			},
		},
	}); err != nil {
		return nil, ConnectionSideUnknown, err
	}

	resp, err := negotiation.Recv()
	if err != nil {
		return nil, ConnectionSideUnknown, err
	}

	side := ConnectionSideUnknown

	switch p := resp.Which.(type) {
	case *pb.NegotiateResponse_Offer_:
		// We are polite, we need to rollback.
		// TODO: Just set the local description to rollback when it's supported.
		peerConn, err = makePeerConn()
		if err != nil {
			return nil, ConnectionSideUnknown, err
		}

		if err := peerConn.SetRemoteDescription(webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: p.Offer.Sdp}); err != nil {
			return nil, ConnectionSideUnknown, err
		}

		answer, err := peerConn.CreateAnswer(nil)
		if err != nil {
			return nil, ConnectionSideUnknown, err
		}
		if err := peerConn.SetLocalDescription(answer); err != nil {
			return nil, ConnectionSideUnknown, err
		}

		if err := negotiation.Send(&pb.NegotiateRequest{
			Which: &pb.NegotiateRequest_Answer_{
				Answer: &pb.NegotiateRequest_Answer{
					Sdp: peerConn.LocalDescription().SDP,
				},
			},
		}); err != nil {
			return nil, ConnectionSideUnknown, err
		}
		side = ConnectionSideAnswerer
	case *pb.NegotiateResponse_Answer_:
		// We are impolite, keep trucking.
		if err := peerConn.SetRemoteDescription(webrtc.SessionDescription{Type: webrtc.SDPTypeAnswer, SDP: p.Answer.Sdp}); err != nil {
			return nil, ConnectionSideUnknown, err
		}
		side = ConnectionSideOfferer
	default:
		return nil, ConnectionSideUnknown, fmt.Errorf("unexpected packet: %v", p)
	}

	allCandidatesSent := make(chan struct{})
	allCandidatesReceived := make(chan struct{})

	go func() {
		for {
			resp, err := negotiation.Recv()
			if err != nil {
				log.Printf("signor: error received during negotiation: %s", err)
				break
			}

			iceCandidateResp, ok := resp.Which.(*pb.NegotiateResponse_IceCandidate)
			if !ok {
				log.Printf("signor: received unexpected response from server: %s", resp)
				continue
			}

			if iceCandidateResp.IceCandidate.IceCandidate == "" {
				log.Printf("signor: all ICE candidates received from peer")
				close(allCandidatesReceived)
				return
			}

			var iceCandidateInit webrtc.ICECandidateInit
			if err := json.Unmarshal([]byte(iceCandidateResp.IceCandidate.IceCandidate), &iceCandidateInit); err != nil {
				log.Printf("signor: failed to unmarshal ICE candidate: %s", err)
				continue
			}

			log.Printf("signor: received ICE candidate from peer: %s", iceCandidateResp.IceCandidate.IceCandidate)

			if err := peerConn.AddICECandidate(iceCandidateInit); err != nil {
				log.Printf("signor: failed to add ICE candidate: %s", err)
				continue
			}
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-allCandidatesSent:
		}
		select {
		case <-ctx.Done():
			return
		case <-allCandidatesReceived:
		}
		negotiation.CloseSend()
	}()

	peerConn.OnConnectionStateChange(func(pcs webrtc.PeerConnectionState) {
		if pcs == webrtc.PeerConnectionStateClosed {
			negotiation.CloseSend()
		}
	})

	peerConn.OnICECandidate(func(iceCandidate *webrtc.ICECandidate) {
		var marshaledICECandidate []byte
		if iceCandidate != nil {
			var err error
			marshaledICECandidate, err = json.Marshal(iceCandidate.ToJSON())
			if err != nil {
				log.Printf("signor: failed to marshal ICE candidate: %s", err)
				return
			}
		}

		if err := negotiation.Send(&pb.NegotiateRequest{
			Which: &pb.NegotiateRequest_IceCandidate{
				IceCandidate: &pb.NegotiateRequest_ICECandidate{
					IceCandidate: string(marshaledICECandidate),
				},
			},
		}); err != nil {
			log.Printf("signor: failed to send ICE candidate: %s", err)
		}

		if iceCandidate == nil {
			log.Printf("signor: all ICE candidates sent to peer")
			close(allCandidatesSent)
		} else {
			log.Printf("signor: sent ICE candidate to peer: %s", string(marshaledICECandidate))
		}
	})
	return peerConn, side, nil
}
