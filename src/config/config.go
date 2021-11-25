package config

type Config struct {
	DB DB
}

func Init() Config {
	return Config{}
}
