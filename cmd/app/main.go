package main

import (
	"context"
	"os"

	"github.com/LiquidCats/go-project-template/configs"
	"github.com/LiquidCats/go-project-template/internal/adapter/repository/database/postgresql"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"

	_ "go.uber.org/automaxprocs"
)

const app = "app"

func main() {
	logger := zerolog.New(os.Stdout)

	zerolog.DefaultContextLogger = &logger

	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	cfg, err := configs.Load(app)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("failed to load config")
	}

	conn, err := pgx.Connect(ctx, cfg.Postgres.ToDSN())
	defer func() {
		if err := conn.Close(ctx); err != nil {
			logger.Fatal().Stack().Err(err).Msg("close connection")
		}
	}()
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("connect to database")
	}

	if err := postgresql.Migrate(conn); err != nil {
		logger.Fatal().Stack().Err(err).Msg("migrate")
	}
}
