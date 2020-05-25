package main

import (
	"context"
	"log"
	"time"

	proto "github.com/PieterVoorwinden/client-server/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Example struct{}

func (e *Example) ListExample(ctx context.Context, req *proto.ExampleRequest, rsp *proto.ExampleResponse) error {
	now := time.Now()
	rsp.Example = &proto.Example{
		CurrentTime: &now,
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.example"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	if err := proto.RegisterExampleServiceHandler(service.Server(), &Example{}); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
