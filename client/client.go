package main

import (
	"context"
	"fmt"
	pb "github.com/axrez/disys-mini-project-2"
	"google.golang.org/grpc"
	"log"
	"time"
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

	fmt.Printf("Participant received with ID: %d and Name: %s \n", p.Id, p.Name)

	go Chat(c, ctx, p, lTime)
	for true {
	}
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

func GetParicipantTextMessage() string {
	var textMessage string
	fmt.Print("Please type a message: ")
	fmt.Scanln(&textMessage)
	return textMessage
}

func PublishMessage(c pb.ChittyChatClient, ctx context.Context, textMessage string, lTime int32) {
	message := &pb.PublishMessage{
		Message: textMessage,
		LTime:   lTime,
	}
	_, err := c.Publish(ctx, message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
}

func LeaveChat(c pb.ChittyChatClient, ctx context.Context, p Participant, lTime int32) {
	message := &pb.LeaveMessage{
		Id:    p.Id,
		LTime: lTime,
	}
	_, err := c.Leave(ctx, message)
	if err != nil {
		log.Fatalf("%s failed to leave chat: %v", p.Name, err)
	}
}

func Chat(c pb.ChittyChatClient, ctx context.Context, p Participant, lTime int32) {
	for true {
		text := GetParicipantTextMessage()
		if text == "\\leave" {
			LeaveChat(c, ctx, p, lTime)
		} else {
			PublishMessage(c, ctx, text, lTime)
		}
	}
}
