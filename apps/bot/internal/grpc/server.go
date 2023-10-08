package grpc

import (
	"context"
	"fmt"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/bot"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type BotServer struct {
	bot.UnimplementedBotServer
}

func New(lc fx.Lifecycle) error {
	server := &BotServer{}

	grpcNetListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatal("cannot create listener")
	}

	grpcServer := grpc.NewServer()
	bot.RegisterBotServer(grpcServer, server)
	reflection.Register(grpcServer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go grpcServer.Serve(grpcNetListener)
			log.Println("Server started on", grpcNetListener.Addr().String())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.Stop()
			return nil
		},
	})

	return nil
}
