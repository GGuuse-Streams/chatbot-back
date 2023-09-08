package main

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/internal/config"
	"github.com/GGuuse-Streams/chatbot-back/internal/db"
	"github.com/GGuuse-Streams/chatbot-back/internal/db/queries"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"strconv"
)

func main() {
	app := fiber.New()

	fx.New(
		fx.Provide(config.New),
		fx.Provide(db.New),
		fx.Provide(queries.New),

		fx.Invoke(func(lifecycle fx.Lifecycle, cfg *config.Config, q *queries.Queries) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					app.Get("/channel", func(c *fiber.Ctx) error {
						channels, err := q.GetChannels(c.Context())
						if err != nil {
							return err
						}

						return c.JSON(channels)
					})

					app.Get("/channel/:id", func(c *fiber.Ctx) error {
						id, err := strconv.Atoi(c.Params("id"))
						if err != nil {
							return err
						}

						channel, err := q.GetChannel(c.Context(), int32(id))
						if err != nil {
							return err
						}

						return c.JSON(channel)
					})

					app.Post("/channel", func(c *fiber.Ctx) error {
						var params queries.CreateChannelParams
						if err := c.BodyParser(&params); err != nil {
							return err
						}

						channel, err := q.CreateChannel(c.Context(), params)
						if err != nil {
							return err
						}

						return c.JSON(channel)
					})

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
		}),
	).Run()
}
