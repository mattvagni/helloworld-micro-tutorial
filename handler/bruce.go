package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	bruce "bruce/proto"
)

type Bruce struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Bruce) Call(ctx context.Context, req *bruce.Request, rsp *bruce.Response) error {
	log.Info("Received Bruce.Call request")
	rsp.Msg = "Hello Mr." + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Bruce) Stream(ctx context.Context, req *bruce.StreamingRequest, stream bruce.Bruce_StreamStream) error {
	log.Infof("Received Bruce.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&bruce.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Bruce) PingPong(ctx context.Context, stream bruce.Bruce_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&bruce.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
