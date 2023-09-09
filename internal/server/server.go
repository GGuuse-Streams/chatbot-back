package server

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/internal/config"
	"github.com/GGuuse-Streams/chatbot-back/internal/server/router"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Backend for GGuuse-Streams chatbot",
	})

	return app
}

func Start(lc fx.Lifecycle, cfg *config.Config, app *fiber.App, router *router.Router) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.Register()

			go func() {
				if err := app.Listen(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := app.Shutdown(); err != nil {
				return err
			}

			return nil
		},
	})
}
