package main

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/clients"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/bot"
)

func main() {
	c := clients.NewBotClient()

	_, err := c.Join(context.Background(), &bot.JoinOrLeaveRequest{
		Channel: "gguuse",
	})
	if err != nil {
		return
	}
}
