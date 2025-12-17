package config

import "os"

type Config struct {
	DBURL string
}

func Load() Config {
	return Config{
		DBURL: os.Getenv("DATABASE_URL"),
	}
}
