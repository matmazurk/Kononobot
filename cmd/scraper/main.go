package main

import (
	"fmt"
	"time"

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

	t := time.Date(2021, time.June, 1, 5, 0, 0, 0, time.UTC)
	mleczny, err := rest.GetListResponse(config.ApiURL, config.ApiKey, MLECZNYID, t)
	if err != nil {
		fmt.Println(err)
	}

	loc, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(mleczny.Items))
	fmt.Println(mleczny.Items[0].Snippet.PublishTime.In(loc).Format(time.RFC3339))
}
