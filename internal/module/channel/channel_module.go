package channel

import (
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/repository"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router struct {
	App fiber.Router

	Service *service.ChannelService
}

var NewChannelModule = fx.Options(
	fx.Provide(repository.NewChannelRepository),
	fx.Provide(service.NewChannelService),

	fx.Provide(NewChannelRouter),
)

func NewChannelRouter(app *fiber.App, service *service.ChannelService) *Router {
	return &Router{
		App:     app,
		Service: service,
	}
}

func (r *Router) RegisterChannelRoutes() {
	r.App.Route("/channel", func(router fiber.Router) {
		router.Get("/", r.Service.GetChannels)
		router.Get("/:id", r.Service.GetChannel)
		router.Post("/", r.Service.CreateChannel)
		router.Delete("/:id", r.Service.DeleteChannel)
		router.Put("/", r.Service.UpdateChannel)
	})
}
