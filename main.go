package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hritikkhetan/go-grpc-time-svc/pkg/config"
	"github.com/hritikkhetan/go-grpc-time-svc/pkg/pb"
	"github.com/hritikkhetan/go-grpc-time-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Time Svc on", c.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterTimeServiceServer(grpcServer, &services.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
