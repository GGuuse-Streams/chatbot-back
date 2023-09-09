package router

import (
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App fiber.Router

	ChannelRouter *channel.Router
}

func New(
	app *fiber.App,
	channelRouter *channel.Router,
) *Router {
	return &Router{
		App:           app,
		ChannelRouter: channelRouter,
	}
}

func (r *Router) Register() {
	r.ChannelRouter.RegisterChannelRoutes()
}
