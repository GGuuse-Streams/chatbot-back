package client

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/bot/internal/client/handlers"
	"github.com/GGuuse-Streams/chatbot-back/libs/config"
	"github.com/GGuuse-Streams/chatbot-back/libs/queries"
	irc "github.com/gempir/go-twitch-irc/v4"
	"go.uber.org/fx"
	"log"
)

type TwitchClient struct {
	*irc.Client

	cfg *config.Config
	q   *queries.Queries
}

func New(lc fx.Lifecycle, cfg *config.Config, q *queries.Queries) *TwitchClient {
	client := &TwitchClient{
		Client: irc.NewClient(cfg.Bot.Username, cfg.Bot.AccessToken),
		q:      q,
		cfg:    cfg,
	}

	client.start(lc)

	return client
}

func (c *TwitchClient) start(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			c.setupHandlers()

			go func() {
				if err := c.Connect(); err != nil {
					log.Fatal("Unable to connect to Twitch IRC")
				}
			}()

			if err := c.initialJoins(); err != nil {
				return err
			}

			return nil
		},
		OnStop: func(_ context.Context) error {
			return c.Client.Disconnect()
		},
	})
}

func (c *TwitchClient) setupHandlers() {
	c.OnConnect(handlers.OnConnect)
	c.OnSelfJoinMessage(handlers.OnSelfJoin)
	c.OnSelfPartMessage(handlers.OnSelfPart)
	c.OnPrivateMessage(handlers.OnPrivateMessage)
}

func (c *TwitchClient) initialJoins() error {
	ctx := context.Background()
	channels, err := c.q.GetChannels(ctx)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		c.Join(channel.TwitchName)
	}

	return nil
}
