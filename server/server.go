package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/xvbnm48/grpc-basic/gen/proto"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	fmt.Println("server start")
	listen, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}

}
