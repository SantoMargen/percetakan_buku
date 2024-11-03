package db

import (
	"context"
	"fmt"
	"log"
	"siap_app/config"

	"github.com/go-redis/redis/v8"
)

func RedisInit(config config.DBConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%v:%v", config.RedisHHost, config.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.RedisPass,
		DB:       0,
		PoolSize: 10,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Coneccted to redis")

	return client, nil
}
