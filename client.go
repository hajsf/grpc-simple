package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/hajsf/grpc-simple/chat"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.MessageRequest{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	res, err := c.BroadcastMessage(context.Background(), &chat.Message{Body: "Message to Broadcast!"})
	if err != nil {
		log.Fatalf("Error when calling Broadcast Message: %s", err)
	}
	log.Printf("Response from server: %s", res.Body)

}
