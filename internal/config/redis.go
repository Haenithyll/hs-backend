package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Publisher struct {
	client  *redis.Client
	channel string
	ctx     context.Context
}

func NewPublisher() (*Publisher, error) {
	redisAddr := GetEnvOrPanic("REDIS_ADDR")
	redisPassword := "" //GetEnvOrPanic("REDIS_PASSWORD")
	redisChannel := GetEnvOrPanic("REDIS_CHANNEL")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Publisher{
		client:  rdb,
		channel: redisChannel,
		ctx:     ctx,
	}, nil
}

func (p *Publisher) Publish(message string) error {
	return p.client.Publish(p.ctx, p.channel, message).Err()
}

func (p *Publisher) Close() error {
	return p.client.Close()
}
