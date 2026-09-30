package config

type Config struct {
	Capacity   int
	RefillRate float64
}

func LoadConfig() Config {

	return Config{
		Capacity:   5,
		RefillRate: 1,
	}
}
