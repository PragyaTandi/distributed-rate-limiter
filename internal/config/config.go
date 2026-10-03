package config

import "os"

type Config struct {
	Capacity   int
	RefillRate float64
	Port       string
}

func LoadConfig() Config {

	return Config{
		Capacity:   5,
		RefillRate: 1,
		Port:       os.Getenv("PORT"),
	}
}
