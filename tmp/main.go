package main

package main

import (
	"context"
	"log"

	"blog-application/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewAuthServiceClient(conn)
	client.Signup(context.Background(), &proto.SignupRequest{Username: "Carl", Email: "carl@gmail.com", Password: "examplestring"})
}