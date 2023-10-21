package main

import (
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/client"
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/db"
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/grpc"
	"github.com/GGuuse-Streams/chatbot-back/libs/config"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/clients"
	"go.uber.org/fx"
	"log"
)

func main() {
	log.SetPrefix("bot microservice: ")
	fx.New(
		fx.NopLogger,

		fx.Provide(config.New),
		fx.Provide(client.NewTwitch),

		fx.Provide(clients.NewCommands),

		db.NewDB,

		fx.Invoke(grpc.New),
	).Run()
}
