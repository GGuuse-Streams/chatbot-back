package handlers

import (
	"github.com/gempir/go-twitch-irc/v4"
	"log"
)

func OnSelfJoin(m twitch.UserJoinMessage) {
	log.Printf("Joined %s", m.Channel)
}
