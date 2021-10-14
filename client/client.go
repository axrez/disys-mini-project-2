package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/axrez/disys-mini-project-2"
	"google.golang.org/grpc"
)

type Participant struct {
	Name string
	Id   int32
}

const address = "localhost:50051"

func main() {
	name := GetParticipantName()
	var lTime int32 = 1

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChittyChatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	p := JoinChat(c, ctx, name, lTime)

	fmt.Printf("Participant received with ID: %d and Name: %s", p.Id, p.Name)
}

func JoinChat(c pb.ChittyChatClient, ctx context.Context, name string, lTime int32) Participant {
	message := &pb.JoinMessage{
		Name:  name,
		LTime: lTime,
	}
	r, err := c.Join(ctx, message)
	if err != nil {
		log.Fatalf("%s failed to join: %v", name, err)
	}
	return Participant{Id: r.GetId(), Name: r.GetName()}
}

func GetParticipantName() string {
	var name string
	fmt.Print("Please type your name: ")
	fmt.Scanln(&name)
	return name
}
