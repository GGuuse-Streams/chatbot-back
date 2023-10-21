package main

import (
	"github.com/GGuuse-Streams/chatbot-back/apps/commands/internal/db"
	"github.com/GGuuse-Streams/chatbot-back/apps/commands/internal/grpc"
	"github.com/GGuuse-Streams/chatbot-back/libs/config"
	"go.uber.org/fx"
	"log"
)

func main() {
	log.SetPrefix("commands microservice: ")
	fx.New(
		fx.NopLogger,

		fx.Provide(config.New),

		db.NewDB,

		fx.Invoke(grpc.New),
	).Run()
}
