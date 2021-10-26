package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)

	name := GetParticipantName(reader)
	var lTime int32 = 1

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChittyChatClient(conn)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	p := JoinChat(c, ctx, name, lTime)

	log.Println("Welcome to chittychat! Write '\\leave' to leave the chat")
	log.Printf("Participant received with ID: %d and Name: %s \n", p.Id, p.Name)

	go Listen(SubscribeChat(c, ctx, p, 1))
	Chat(c, ctx, p, lTime, reader)
}

func GetParticipantName(r *bufio.Reader) string {
	fmt.Print("Please type your name: ")
	rawName, _ := r.ReadString('\n')

	name := strings.TrimSpace(rawName)

	if len(name) == 0 {
		fmt.Println("The name must be more than 0 characters")
		return GetParticipantName(r)
	}

	return name
}

func GetParicipantTextMessage(r *bufio.Reader) string {
	fmt.Print("Please type a message: ")

	rawMessage, _ := r.ReadString('\n')
	message := strings.TrimSpace(rawMessage)

	return message
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

func SubscribeChat(c pb.ChittyChatClient, ctx context.Context, p Participant, lTime int32) pb.ChittyChat_SubscribeClient {
	message := &pb.SubscribeMessage{
		Id:    p.Id,
		LTime: lTime,
	}
	stream, err := c.Subscribe(ctx, message)
	if err != nil {
		log.Fatalf("%s failed to subscribe to chat: %v", p.Name, err)
	}

	return stream

}

func PublishMessage(c pb.ChittyChatClient, ctx context.Context, textMessage string, p Participant, lTime int32) {
	message := &pb.PublishMessage{
		Message: textMessage,
		Id:      p.Id,
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
	os.Exit(0)
}

func Chat(c pb.ChittyChatClient, ctx context.Context, p Participant, lTime int32, r *bufio.Reader) {
	for {
		text := GetParicipantTextMessage(r)
		if text == "\\leave" {
			LeaveChat(c, ctx, p, lTime)
		} else {
			PublishMessage(c, ctx, text, p, lTime)
		}
	}
}

func Listen(stream pb.ChittyChat_SubscribeClient) {
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			time.Sleep(1 * time.Second)
		} else if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}
		if message != nil {
			log.Println(message.Message)
		}
	}
}
