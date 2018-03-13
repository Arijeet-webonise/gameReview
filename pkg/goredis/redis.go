package goredis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
}

type IRedisClient interface {
	Get(string) (string, error)
	Set(string, interface{}, time.Duration) error
	Del(...string) (int64, error)
	Exists(...string) (int64, error)
	Close() error
}

func NewClient(config map[string]string) (*RedisClient, error) {
	rc := new(RedisClient)

	db, err := strconv.Atoi(config["DB"])
	if err != nil {
		return rc, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config["Addr"],
		Password: config["Password"],
		DB:       db,
	})

	rc.Client = client

	rc.Client = client
	return rc, nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.Client.Get(key).Result()
}

func (rc *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return rc.Client.Set(key, value, expiration).Err()
}

func (rc *RedisClient) Del(keys ...string) (int64, error) {
	return rc.Client.Del(keys...).Result()
}

func (rc *RedisClient) Exists(keys ...string) (int64, error) {
	return rc.Client.Exists(keys...).Result()
}

func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}
