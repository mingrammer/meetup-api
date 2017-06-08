package config

import (
	"os"
)

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
	ClientID     string
	ClientSecret string
	TokenURL     string
}

func GetConfig() *Config {
	DefaultHost := "localhost"
	DefaultPort := 8080

	return &Config{
		Secret: os.Getenv("API_SECRET_VALUE"),
		Host:   DefaultHost,
		Port:   DefaultPort,
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     "meetup",
			Charset:  "utf8",
		},
		SlackApp: &SlackAppConfig{
			ClientID:     os.Getenv("SLACK_CLIENT_ID"),
			ClientSecret: os.Getenv("SLACK_CLIENT_SECRET"),
			TokenURL:     "https://slack.com/api/oauth.access",
		},
	}
}
