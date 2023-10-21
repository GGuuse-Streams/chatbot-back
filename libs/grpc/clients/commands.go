package clients

import (
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/commands"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCommands() commands.CommandsClient {
	conn, err := grpc.Dial("commands:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return commands.NewCommandsClient(conn)
}
