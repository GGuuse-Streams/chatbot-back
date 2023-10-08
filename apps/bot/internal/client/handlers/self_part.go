package handlers

import (
	"github.com/gempir/go-twitch-irc/v4"
	"log"
)

func OnSelfPart(m twitch.UserPartMessage) {
	log.Printf("Left %s", m.Channel)
}
