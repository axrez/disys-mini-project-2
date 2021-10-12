package main

import (
	"context"
	"fmt"
	"log"
	pb "github.com/axrez/disys-mini-project-2"
)

type Participant struct {
	Name string
	Id   int32
}

func main() {
}

func JoinChat (c pb.ChittyChatClient, ctx context.Context, name string, lTime int32) Participant{
	message := &pb.JoinMessage{
		Name: name,
		LTime: lTime,
	}
	_, err := c.Join(ctx, message)
	if err != nil {
		log.Fatalf("%s failed to join: %v", name, err)
	}
	return Participant{Id: 1, Name: name}
}

func Welcome() string {
	var name string
	fmt.Print("Please type your name: ")
	fmt.Scanln(&name)
	return name
}
