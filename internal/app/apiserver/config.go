package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:database_url`
	SessionKey  string `toml:session_key'`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		DatabaseURL: "host=localhost dbname=restapi_dev sslmode=disable",
		SessionKey:  "2c70e12b7a0646f92279f427c7b38e7334d8e5389cff167a1dc30e73f826b683",
	}
}
