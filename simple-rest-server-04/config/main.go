package config

type Config struct {
	Version string
}

func NewConfig() Config {
	return Config{"1.0.0"}
}
