package main

import (
	"context"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/matmazurk/Kononobot/internal/handlers"
	"github.com/matmazurk/Kononobot/internal/services"
	"github.com/matmazurk/Kononobot/pkg/persistence"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	env "github.com/Netflix/go-env"
)

type Config struct {
	ApiKey string `env:"API_KEY"`
	DbDns  string `env:"DB_DNS"`
}

func main() {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot unmarshal config from env")
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("kononobot starting...")

	db := connectToPostgresDB(context.Background(), cfg.DbDns)
	pgClient := persistence.NewPostgresClient(db)
	ytHandler := handlers.NewYT(cfg.ApiKey)
	kbotService := services.NewKBot(ytHandler, pgClient)
	err = kbotService.Serve()
	if err != nil {
		log.Error().Err(err).Msg("kononobot service returned with error")
	}
}

func connectToPostgresDB(ctx context.Context, dns string) *sqlx.DB {
	db, err := sqlx.Open("postgres", dns)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
		log.Info().Int("ping_counter", i).Msg("pinging database...")
	}

	if err != nil {
		log.Fatal().Err(err).Msg("didn't connect to db")
	}
	log.Info().Msg("postgres connection established")
	return db
}
