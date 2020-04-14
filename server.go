package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "gorpc/proto"
	"log"
	"net"
	"time"
)

type Server struct {
	pb.UnimplementedDemoServer
}

func NewServerStart() error {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %s", err)
	}

	serv := &Server{}
	grpcServ := grpc.NewServer()

	fmt.Println("server started ...")
	pb.RegisterDemoServer(grpcServ, serv)

	return grpcServ.Serve(lis)
}

func (serv *Server) ServerVersion(ctx context.Context, in *pb.EmptyRequest) (*pb.VersionReply, error) {
	log.Printf("ServerVersion requested")
	resp := pb.VersionReply{
		Version: "gRPC Server 1.0.0",
	}
	return &resp, nil
}

func (serv *Server) Watch(in *pb.VideoRequest, s pb.Demo_WatchServer) error {
	for i := 0; i < 5; i++ {
		reply := &pb.VideoReply{
			Block: []byte(fmt.Sprintf("video: %s, frame: %d", in.Id, i)),
		}
		if err := s.Send(reply); err != nil {
			return err
		}
		<-time.After(time.Second * 1)
	}
	return nil
}
