package client

import (
	"context"
	"net/http"

	"github.com/pion/webrtc/v3"
	"github.com/murkland/signor/pb"
)

type Client struct {
	rpcClient pb.SessionService
}

func New(address string) *Client {
	return &Client{pb.NewSessionServiceProtobufClient(address, http.DefaultClient)}
}

func (c *Client) Offer(ctx context.Context, sessionID []byte, peerConn *webrtc.PeerConnection) error {
	gatherComplete := webrtc.GatheringCompletePromise(peerConn)
	offer, err := peerConn.CreateOffer(nil)
	if err != nil {
		return err
	}
	if err = peerConn.SetLocalDescription(offer); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-gatherComplete:
	}

	answer, err := c.rpcClient.Offer(ctx, &pb.OfferRequest{
		SessionId:  sessionID,
		MyOfferSdp: peerConn.LocalDescription().SDP,
	})
	if err != nil {
		return err
	}

	if err = peerConn.SetRemoteDescription(webrtc.SessionDescription{Type: webrtc.SDPTypeAnswer, SDP: answer.TheirAnswerSdp}); err != nil {
		return err
	}

	return nil
}

func (c *Client) Answer(ctx context.Context, sessionID []byte, peerConn *webrtc.PeerConnection) error {
	offer, err := c.rpcClient.GetOffer(ctx, &pb.GetOfferRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return err
	}

	if err := peerConn.SetRemoteDescription(webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: offer.TheirOfferSdp}); err != nil {
		return err
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConn)
	answer, err := peerConn.CreateAnswer(nil)
	if err != nil {
		return err
	}
	if err := peerConn.SetLocalDescription(answer); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-gatherComplete:
	}

	if _, err := c.rpcClient.Answer(ctx, &pb.AnswerRequest{
		SessionId:   sessionID,
		MyAnswerSdp: peerConn.LocalDescription().SDP,
	}); err != nil {
		return err
	}

	return nil
}
