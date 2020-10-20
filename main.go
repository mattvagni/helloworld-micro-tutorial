package main

import (
	"bruce/handler"
	pb "bruce/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("bruce"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterBruceHandler(srv.Server(), new(handler.Bruce))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
