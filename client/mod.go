package client

import (
	"context"
	"fmt"

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

	gatherComplete := webrtc.GatheringCompletePromise(peerConn)
	offer, err := peerConn.CreateOffer(nil)
	if err != nil {
		return nil, ConnectionSideUnknown, err
	}
	if err = peerConn.SetLocalDescription(offer); err != nil {
		return nil, ConnectionSideUnknown, err
	}

	select {
	case <-ctx.Done():
		return nil, ConnectionSideUnknown, ctx.Err()
	case <-gatherComplete:
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

	switch p := resp.Which.(type) {
	case *pb.NegotiateResponse_Offer_:
		// We are polite, we need to rollback.
		// TODO: Just set the local description to rollback when it's supported.
		peerConn, err := makePeerConn()
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
		return peerConn, ConnectionSideAnswerer, nil
	case *pb.NegotiateResponse_Answer_:
		// We are impolite, keep trucking.
		if err := peerConn.SetRemoteDescription(webrtc.SessionDescription{Type: webrtc.SDPTypeAnswer, SDP: p.Answer.Sdp}); err != nil {
			return nil, ConnectionSideUnknown, err
		}
		return peerConn, ConnectionSideOfferer, nil
	default:
		return nil, ConnectionSideUnknown, fmt.Errorf("unexpected packet: %v", p)
	}
}
