package main

import (
	"context"
	"log"
	"time"

	proto "github.com/PieterVoorwinden/client-server/proto"
	"github.com/micro/go-micro/v2"
	client "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	time.Sleep(2 * time.Second)
	srv := proto.NewExampleService("go.micro.srv.example", client.NewClient())
	rsp, err := srv.ListExample(context.Background(), &proto.ExampleRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(rsp.Example.CurrentTime)
}
