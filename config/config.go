package config

import (
	"os"
)

type DBConfig struct {
	Host          string
	Port          int
	User          string
	Password      string
	DBName        string
	SSLMode       string
	SecretDecrypt string
	RedisHHost    string
	RedisPass     string
	RedisPort     string
}

func LoadDBConfig() *DBConfig {

	return &DBConfig{
		Host:          os.Getenv("DB_HOST"),
		Port:          5432,
		User:          os.Getenv("DB_USER"),
		Password:      os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		SSLMode:       os.Getenv("SSL_MODE"),
		SecretDecrypt: os.Getenv("DECRYPT_KEY"),
		RedisHHost:    os.Getenv("REDIS_HOST"),
		RedisPass:     os.Getenv("REDIS_PASS"),
		RedisPort:     os.Getenv("REDIS_PORT"),
	}
}
