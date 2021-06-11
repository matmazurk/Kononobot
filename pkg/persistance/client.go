package persistance

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Repository interface {
	InsertFilm(ctx context.Context, film Film) error
	GetLatestFilm(ctx context.Context, channelID string) (Film, error)
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

func (c client) InsertFilm(ctx context.Context, film Film) error {
	key := film.ID
	value, err := json.Marshal(film)
	if err != nil {
		return err
	}
	err = c.Set(ctx, key, value, 0).Err()
	return err
}

func (c client) GetLatestFilm(ctx context.Context, channelID string) (Film, error) {
	return Film{}, nil
}
