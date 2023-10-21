package handlers

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/libs/grpc/generated/commands"
	"github.com/gempir/go-twitch-irc/v4"
	"log"
	"strings"
)

func OnPrivateMessage(tc *twitch.Client, cc commands.CommandsClient) func(m twitch.PrivateMessage) {
	return func(m twitch.PrivateMessage) {
		if !strings.HasPrefix(m.Message, "!") {
			return
		}

		answer, err := cc.ProcessCommand(context.Background(), &commands.ProcessCommandRequest{
			Command: m.Message,
			Channel: m.Channel,
		})
		if err != nil {
			log.Println(err)
			return
		}

		tc.Say(m.Channel, answer.Response)

		log.Printf("%s said %s", m.User.DisplayName, m.Message)
	}
}

//func OnPrivateMessage(m twitch.PrivateMessage) {
//	if !strings.HasPrefix(m.Message, "!") {
//		return
//	}
//
//	//answer, err := h.commands.ProcessCommand(context.Background(), &commands.ProcessCommandRequest{
//	//	Command: m.Message,
//	//	Channel: m.Channel,
//	//})
//	//if err != nil {
//	//	log.Println(err)
//	//	return
//	//}
//	//
//	//asd := client.NewTwitch(nil, nil, nil, nil)
//
//	log.Printf("%s said %s", m.User.DisplayName, m.Message)
//}
