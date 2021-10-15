package main

import (
	"context"
	"log"
	"net"

	pb "github.com/axrez/disys-mini-project-2"
	"google.golang.org/grpc"

	utils "github.com/axrez/disys-mini-project-2/utils"
)

const port = ":50051"

type streamWrapper struct {
	messageStream grpc.ServerStream
	error chan error
} 

type participant struct {
	name          string
	streamWrap *streamWrapper
}

type server struct {
	pb.UnimplementedChittyChatServer
	lTime        int32
	participants map[int]participant
}

func (s *server) Join(ctx context.Context, in *pb.JoinMessage) (*pb.JoinReplyMessage, error) {
	utils.CalcNextLTime(&s.lTime, &in.LTime)

	newId := len(s.participants)

	s.participants[newId] = participant{
		name:          in.GetName(),
		streamWrap: &streamWrapper{
			messageStream: nil,
			error: nil,
		},
	}

	return &pb.JoinReplyMessage{Id: int32(newId), Name: in.GetName(), LTime: s.lTime}, nil
}

func (s *server) Subscribe(in *pb.SubscribeMessage, stream pb.ChittyChat_SubscribeServer) error {
	streamWrap := streamWrapper {
		messageStream: stream,
			error: make(chan error),
		}
	part := s.participants[int(in.GetId())]
	part.streamWrap = &streamWrap
	s.participants[int(in.GetId())] = part
	return <- streamWrap.error
}

func (s *server) Publish(ctx context.Context, in *pb.PublishMessage) (*pb.EmptyReturn, error) {
	for _, part := range s.participants {
		if part.streamWrap.messageStream != nil {
			part.streamWrap.messageStream.SendMsg(&pb.BroadcastMessage{Message: in.GetMessage(), LTime: 0})
		}
	}

	return &pb.EmptyReturn{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterChittyChatServer(s, &server{lTime: 0, participants: make(map[int]participant)})

	log.Printf("server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
