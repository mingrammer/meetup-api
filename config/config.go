package config

import "fmt"

type Config struct {
	Secret   string
	Host     string
	Port     int
	DB       *DBConfig
	SlackApp *SlackAppConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

type SlackAppConfig struct {
	ClientId     string
	ClientSecret string
	RedirectURL  string
}

func GetConfig() *Config {
	DefaultHost := "localhost"
	DefaultPort := 8080

	return &Config{
		Secret: "...",
		Host:   DefaultHost,
		Port:   DefaultPort,
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "...",
			Password: "...",
			Name:     "meetup",
			Charset:  "utf8",
		},
		SlackApp: &SlackAppConfig{
			ClientId:     "...",
			ClientSecret: "...",
			RedirectURL:  fmt.Sprintf("http://%s:%d/auth", DefaultHost, DefaultPort),
		},
	}
}
