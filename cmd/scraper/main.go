package main

import (
	"fmt"

	env "github.com/Netflix/go-env"
)

type Config struct {
	ApiURL string `env:"API_URL"`
	ApiKey string `env:"API_KEY"`
}

func main() {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
