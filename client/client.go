package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/xvbnm48/grpc-basic/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8800", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{
		Msg: "hello everyone",
	})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.Msg)
}
