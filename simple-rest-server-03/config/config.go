package config

// Config is storing app configuration
type Config struct {
	Version string
}

// NewConfig is factory for creating config
func NewConfig() Config {
	return Config{"1.0.0"}
}
