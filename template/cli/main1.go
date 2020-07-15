package main

import (
	"context"
	"fmt"

	//proto "../srv/proto/srv" //这里写你的proto文件放置路劲
	proto "github.com/ke00ke/examples/template/srv/proto/srv"

	//proto "proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewSrvService("go.micro.service.srv", service.Client())

	// Call the greeter
	//rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	rsp, err := greeter.Call(context.TODO(), &proto.Request{Name: "ckh"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)
}
