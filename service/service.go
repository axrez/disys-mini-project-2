package main

import (
	"context"
	"log"
	"net"

	pb "github.com/axrez/mini-project-2"
	"google.golang.org/grpc"

	utils "github.com/axrez/mini-project-2/utils"
)

const port = ":50051"

type server struct {
	pb.UnimplementedChittyChatServer
	lTime int32
}

func (s *server) Join(ctx context.Context, in *pb.JoinMessage) (*pb.JoinReplyMessage, error) {
	utils.CalcNextLTime(&s.lTime, &in.LTime)
	return &pb.JoinReplyMessage{Id: 1, Name: "HardCoded", LTime: s.lTime}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChittyChatServer(s, &server{lTime: 0})
	log.Printf("server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
