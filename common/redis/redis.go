package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client redis.Client

func NewClient(address string, password string, db int) (Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping(context.Background()).Result()
	return Client(*client), err
}

func (client *Client) ObtainLock(key string, expiration time.Duration) error {
	val, err := client.SetNX(context.Background(), key, 1, expiration).Result()
	if err != nil {
	}
	if !val {
		return errors.New("lock is been taken")
	}
	return nil
}

func (client *Client) ReleaseLock(key string) {
	client.Del(context.Background(), key)
}

func (client *Client) Publish(ctx context.Context, channel string, message interface{}) error {
	c := redis.Client(*client)
	err := c.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	c := redis.Client(*client)
	return c.Subscribe(ctx, channels...)
}

func (client *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	c := redis.Client(*client)
	c.Set(ctx, key, value, expiration)
}

func (client *Client) Get(ctx context.Context, key string) *redis.StringCmd {
	c := redis.Client(*client)
	return c.Get(ctx, key)
}

func (client *Client) Del(ctx context.Context, key string) {
	c := redis.Client(*client)
	c.Del(ctx, key)
}
