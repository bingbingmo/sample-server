package main

import (
	"github.com/bingbingmo/sample-server/handler"
	sample "github.com/bingbingmo/sample-server/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("sample"),
	)

	// initialise flags
	service.Init()

	h, err := handler.NewSampleServerHandler()
	if err != nil {
		logger.Fatal("Could not create new handler.")
	}

	err = sample.RegisterSampleServerHandler(service.Server(), h)
	if err != nil {
		logger.Fatalf("Could not register handler.")
	}

	// start the service
	if err := service.Run(); err != nil {
		logger.Fatalf("Could not run sample-server service.")
	}
}
