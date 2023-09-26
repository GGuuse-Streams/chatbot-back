package channel

import (
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/controller"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/repository"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router struct {
	App fiber.Router

	Controller *controller.ChannelController
}

var NewChannelModule = fx.Options(
	fx.Provide(repository.NewChannelRepository),
	fx.Provide(service.NewChannelService),
	fx.Provide(controller.NewChannelController),

	fx.Provide(NewChannelRouter),
)

func NewChannelRouter(app *fiber.App, controller *controller.ChannelController) *Router {
	return &Router{
		App:        app,
		Controller: controller,
	}
}

func (r *Router) RegisterChannelRoutes() {
	r.App.Route("/channel", func(router fiber.Router) {
		router.Get("/", r.Controller.GetChannels)
		router.Get("/:id", r.Controller.GetChannel)
		router.Post("/", r.Controller.CreateChannel)
		router.Delete("/:id", r.Controller.DeleteChannel)
		router.Put("/", r.Controller.UpdateChannel)
	})
}
