package grpc

import (
	"context"
	"fmt"
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/client"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/bot"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type BotServer struct {
	bot.UnimplementedBotServer

	twitch *client.TwitchClient
}

func New(lc fx.Lifecycle, c *client.TwitchClient) error {
	server := &BotServer{
		twitch: c,
	}

	grpcNetListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatal("cannot create listener")
	}

	grpcServer := grpc.NewServer()
	bot.RegisterBotServer(grpcServer, server)
	reflection.Register(grpcServer)

	server.start(lc, grpcServer, grpcNetListener)

	return nil
}

func (b BotServer) start(lc fx.Lifecycle, grpcServer *grpc.Server, grpcNetListener net.Listener) {
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
}

func (b BotServer) Join(_ context.Context, request *bot.JoinOrLeaveRequest) (*emptypb.Empty, error) {
	b.twitch.Join(request.Channel)
	return &emptypb.Empty{}, nil
}

func (b BotServer) Leave(_ context.Context, request *bot.JoinOrLeaveRequest) (*emptypb.Empty, error) {
	b.twitch.Depart(request.Channel)
	return &emptypb.Empty{}, nil
}
