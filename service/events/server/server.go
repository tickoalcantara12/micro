package server

import (
	pb "github.com/tickoalcantara12/micro/v3/proto/events"
	"github.com/tickoalcantara12/micro/v3/service"
	"github.com/tickoalcantara12/micro/v3/service/events/handler"
	"github.com/tickoalcantara12/micro/v3/service/logger"
	"github.com/urfave/cli/v2"
)

// Run the micro broker
func Run(ctx *cli.Context) error {
	// new service
	srv := service.New(
		service.Name("events"),
	)

	// register the handlers
	pb.RegisterStreamHandler(srv.Server(), new(handler.Stream))
	pb.RegisterStoreHandler(srv.Server(), new(handler.Store))

	// run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	return nil
}
