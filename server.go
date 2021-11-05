package main

import (
	"fmt"
	"log"
	"net"

	"github.com/PyBastian/Distribuidos_sg/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Server Iniciado")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Println("Server Final main")
}
