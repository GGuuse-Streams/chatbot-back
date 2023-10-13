package clients

import (
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/bot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewBotClient() bot.BotClient {
	conn, err := grpc.Dial("bot:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := bot.NewBotClient(conn)

	return client
}
