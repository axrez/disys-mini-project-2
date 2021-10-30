package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	pb "github.com/axrez/disys-mini-project-2"
	"google.golang.org/grpc"


	utils "github.com/axrez/disys-mini-project-2/utils"
)

const address = "localhost:50051"

var lTime []int32
var id int32

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := GetParticipantName(reader)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChittyChatClient(conn)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id, lTime = JoinChat(c, ctx, name)
	utils.IncrementLTime(id, &lTime)

	log.Println("Welcome to chittychat! Write '\\leave' to leave the chat")
	log.Printf("ID received: %d\n", id)

	SetupCloseHandler(c, ctx)
	go Listen(SubscribeChat(c, ctx))
	Chat(c, ctx, reader)
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
	rawMessage, _ := r.ReadString('\n')
	message := strings.TrimSpace(rawMessage)

	return message
}

func JoinChat(c pb.ChittyChatClient, ctx context.Context, name string) (int32, []int32) {
	message := &pb.JoinMessage{
		Name:  name,
	}
	r, err := c.Join(ctx, message)
	if err != nil {
		log.Fatalf("%s failed to join: %v", name, err)
	}
	return r.GetId(), r.GetLTime()
}

func SubscribeChat(c pb.ChittyChatClient, ctx context.Context) pb.ChittyChat_SubscribeClient {
	utils.IncrementLTime(id, &lTime)	
	message := &pb.SubscribeMessage{
		Id:    id,
		LTime: lTime,
	}
	stream, err := c.Subscribe(ctx, message)
	if err != nil {
		log.Fatalf("Participant with ID: %d failed to subscribe to chat: %v", id, err)
	}

	return stream
}

func PublishMessage(c pb.ChittyChatClient, ctx context.Context, textMessage string) {
	utils.IncrementLTime(id, &lTime)	
	message := &pb.PublishMessage{
		Message: textMessage,
		Id:      id,
		LTime:   lTime,
	}
	r, err := c.Publish(ctx, message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	utils.CalcNextLTime(id, &lTime, &r.LTime)
}

func LeaveChat(c pb.ChittyChatClient, ctx context.Context) {
	utils.IncrementLTime(id, &lTime)	
	message := &pb.LeaveMessage{
		Id:    id,
		LTime: lTime,
	}
	r, err := c.Leave(ctx, message)
	if err != nil {
		log.Fatalf("Participant with ID: %d failed failed to leave chat:: %v", id, err)
	}

	utils.CalcNextLTime(id, &lTime, &r.LTime)
	os.Exit(0)
}

func Chat(c pb.ChittyChatClient, ctx context.Context, r *bufio.Reader) {
	for {
		text := GetParicipantTextMessage(r)
		if text == "\\leave" {
			LeaveChat(c, ctx)
		} else {
			PublishMessage(c, ctx, text)
		}
	}
}

func Listen(stream pb.ChittyChat_SubscribeClient) {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			time.Sleep(1 * time.Second)
		} else if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}
		if r != nil {
			utils.CalcNextLTime(id, &lTime, &r.LTime)
			log.Println(r.Message + utils.LTimeToString(lTime))
		}
	}
}

func SetupCloseHandler(c pb.ChittyChatClient, ctx context.Context) {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		LeaveChat(c, ctx)
	}()
}
