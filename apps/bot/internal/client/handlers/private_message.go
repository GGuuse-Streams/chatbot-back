package handlers

import (
	"github.com/gempir/go-twitch-irc/v4"
	"log"
)

func OnPrivateMessage(m twitch.PrivateMessage) {
	log.Printf("%s said %s", m.User.DisplayName, m.Message)
}
