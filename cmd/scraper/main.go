package main

import (
	"fmt"

	env "github.com/Netflix/go-env"
	"github.com/matmazurk/Kononobot/internal/scraper"
	"github.com/matmazurk/Kononobot/pkg/persistance"
)

type Config struct {
	ApiURL    string `env:"API_URL"`
	ApiKey    string `env:"API_KEY"`
	RedisURL  string `env:"REDIS_URL"`
	RedisPort int    `env:"REDIS_PORT"`
}

func main() {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := persistance.NewClient(config.RedisURL, "")
	scraper := scraper.NewService(scraper.Config{config.ApiURL, config.ApiKey}, db)
	scraper.Run()
}
