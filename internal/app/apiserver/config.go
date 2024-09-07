package apiserver

type Config struct {
	Address  string `toml:"address"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Address:  "localhost:8080",
		LogLevel: "debug",
	}
}
