package persistance

import (
	"github.com/go-redis/redis/v8"
)

type Repository interface {
	InsertFilm(film Film)
	GetLatestFilm(channelID string) Film
}

type client struct {
	*redis.Client
}

func NewClient(address, password string) client {
	cli := client{}
	cli.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return cli
}

func (c client) InsertFilm(film Film) {

}

func (c client) GetLatestFilm(channelID string) Film {
	return Film{}
}
