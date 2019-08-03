package config

var config *Config

// Config is storing app configuration
type Config struct {
	Version string
}

// NewConfig is factory for creating config
func NewConfig() *Config {
	if config != nil {
		return config
	}

	config = &Config{"1.0.0"}
	return config
}
