package main

import (
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/client"
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/db"
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/grpc"
	"github.com/GGuuse-Streams/chatbot-back/libs/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(client.New),

		db.NewDB,

		fx.Invoke(grpc.New),
	).Run()
}
