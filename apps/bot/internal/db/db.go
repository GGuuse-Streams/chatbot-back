package db

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/libs/config"
	"github.com/GGuuse-Streams/chatbot-back/libs/queries"
	"go.uber.org/fx"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(lc fx.Lifecycle, cfg *config.Config) queries.DBTX {
	conn, err := pgx.Connect(context.Background(), cfg.App.ConnectionString)
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			return conn.Close(context.Background())
		},
	})

	return conn
}

var NewDB = fx.Options(
	fx.Provide(New),
	fx.Provide(queries.New),
)
