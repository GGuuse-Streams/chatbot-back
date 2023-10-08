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
		//fx.Invoke(func(lc fx.Lifecycle, s *grpc.BotServer) { s.Start(lc) }),
		//fx.Invoke(func(lc fx.Lifecycle, c *client.TwitchClient) { c.Start(lc) }),
	).Run()
}
