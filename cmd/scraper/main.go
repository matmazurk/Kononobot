package main

import (
	"fmt"

	env "github.com/Netflix/go-env"
	"github.com/matmazurk/Kononobot/internal/rest"
)

type Config struct {
	ApiURL string `env:"API_URL"`
	ApiKey string `env:"API_KEY"`
}

const (
	MLECZNYID        = "UCPhnxJMplQz9A1DstzRDlEA"
	WOJTEKZBOMBASUID = "UCa3y6lbZByfJ2tSPZ9XrjRw"
)

func main() {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	mleczny, err := rest.GetListResponse(config.ApiURL, config.ApiKey, MLECZNYID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mleczny.Items[0].Snippet.ChannelTitle)
}
