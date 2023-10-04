package db

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/internal/config"
	"github.com/GGuuse-Streams/chatbot-back/internal/db/queries"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/fx"
)

func New(cfg *config.Config) queries.DBTX {
	conn, err := pgx.Connect(context.Background(), cfg.Database.ConnectionString)

	if err != nil {
		panic(err)
	}

	return conn
}

var NewDBModule = fx.Options(
	fx.Provide(New),
	fx.Provide(queries.New),
)
