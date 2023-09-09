package main

import (
	"github.com/GGuuse-Streams/chatbot-back/internal/config"
	"github.com/GGuuse-Streams/chatbot-back/internal/db"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel"
	"github.com/GGuuse-Streams/chatbot-back/internal/server"
	"github.com/GGuuse-Streams/chatbot-back/internal/server/router"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(router.New),
		fx.Provide(server.New),

		db.NewDBModule,

		channel.NewChannelModule,

		fx.Invoke(server.Start),
	).Run()
}
