package config

type Config struct {
	Port string
	Addr string
}

func New() *Config {
	return &Config{
		Port: "8080",
		Addr: ":",
	}
}