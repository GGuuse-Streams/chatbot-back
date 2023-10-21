package grpc

import (
	"context"
	"fmt"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/commands"
	"github.com/GGuuse-Streams/chatbot-back/libs/queries"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
)

type CommandsServer struct {
	commands.UnimplementedCommandsServer

	q *queries.Queries
}

func New(lc fx.Lifecycle, q *queries.Queries) error {
	server := &CommandsServer{
		q: q,
	}

	grpcNetListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 9091))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	commands.RegisterCommandsServer(grpcServer, server)
	reflection.Register(grpcServer)

	server.start(lc, grpcServer, grpcNetListener)

	return nil
}

func (c CommandsServer) start(lc fx.Lifecycle, grpcServer *grpc.Server, grpcNetListener net.Listener) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go grpcServer.Serve(grpcNetListener)
			log.Println("Server started on", grpcNetListener.Addr().String())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})
}

func (c CommandsServer) ProcessCommand(ctx context.Context, req *commands.ProcessCommandRequest) (*commands.ProcessCommandResponse, error) {
	channelId, err := c.q.GetChannelIdByName(ctx, req.Channel)
	if err != nil {
		return nil, err
	}

	// ! TODO: Add command args validation
	if strings.Contains(req.Command, " ") {
		return nil, fmt.Errorf("command cannot contain spaces in beta")
	}

	answer, err := c.q.GetAnswer(ctx, queries.GetAnswerParams{
		ChannelID: channelId,
		Command:   req.Command,
	})
	if err != nil {
		return nil, err
	}

	return &commands.ProcessCommandResponse{
		Response: answer,
	}, nil
}
