package main

import (
	"github.com/ke00ke/examples/template/srv/handler"
	"github.com/ke00ke/examples/template/srv/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	srv "github.com/ke00ke/examples/template/srv/proto/srv"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	srv.RegisterSrvHandler(service.Server(), new(handler.Srv))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.srv", service.Server(), new(subscriber.Srv))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
