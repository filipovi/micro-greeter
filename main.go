package main

import (
	"context"
	"log"

	proto "github.com/filipovi/micro-greeter/proto"
	micro "github.com/micro/go-micro"
)

// Greeter is a struct containing the name
type Greeter struct{}

// Hello send the name preceded with Hello
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	log.Println("Starting...")
	defer log.Println("...bye bye!")

	service := micro.NewService(
		micro.Name("greeter"),
	)

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
