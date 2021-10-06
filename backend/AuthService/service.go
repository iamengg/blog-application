package main

import (
	"context"
	"log"
	"net"

	"blog-application/proto"
	"google.golang.org/grpc"
)

type authServer struct{}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	return &authServer{}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})
	listner, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}
	server.Serve(listner)
}
