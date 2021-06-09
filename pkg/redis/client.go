package redis

import (
	rd "github.com/go-redis/redis/v8"
)

type Repository interface {
}

type client struct {
	*rd.Client
}

func NewClient(address, password string) client {
	cli := client{}
	cli.Client = rd.NewClient(&rd.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return cli
}
