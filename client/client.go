package main

import (
	"context"
	"fmt"

	proto "github.com/filipovi/micro-greeter/proto"
	micro "github.com/micro/go-micro"
)

// Greeter is a struct containing the name
type Greeter struct{}

func main() {
	service := micro.NewService()
	greeter := proto.NewGreeterService("greeter", service.Client())

	// request the Hello method on the Greeter handler
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "Pascal",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
