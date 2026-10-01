package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Username string
	Password string
	Host     string
	Port     string
}

func NewRedisClient(username, password, host, port string) *RedisClient {
	return &RedisClient{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (rc *RedisClient) Connect() *redis.Client {
	// validasi config
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rc.Host, rc.Port),
		Username: rc.Username,
		Password: rc.Password,
	})
}
