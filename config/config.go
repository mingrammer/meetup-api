package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type WebAPIEnv struct {
	Secret   string `envconfig:"API_SECRET_VALUE"`
	Port     int `default:"8080"`
	SlackApp *SlackAppEnv
}

type BotAPIEnv struct {
	Secret          string `envconfig:"API_SECRET_VALUE"`
	Port            int `default:"8081"`
	GoogleAPIKey    string `envconfig:"GOOGLE_API_KEY"`
	GoogleMapAPIKey string `envconfig:"GOOGLE_MAP_API_KEY"`
	WebEndpoint     string `envconfig:"WEB_ENDPOINT"`
	SlackBot        *SlackBotEnv
}

type DBEnv struct {
	Dialect  string `default:"mysql"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `default:"meetup"`
	Charset  string `default:"utf8"`
}

type SlackBotEnv struct {
	BotToken          string `envconfig:"SLACK_BOT_TOKEN"`
	VerificationToken string `envconfig:"SLACK_VERIFICATION_TOKEN"`
	BotID             string `envconfig:"SLACK_BOT_ID" default="meetup"`
	ChannelID         string `envconfig:"SLACK_BOT_CHANNEL_ID"`
}

type SlackAppEnv struct {
	ClientID     string `envconfig:"SLACK_CLIENT_ID"`
	ClientSecret string `envconfig:"SLACK_CLIENT_SECRET"`
	TokenURL     string `default:"https://slack.com/api/oauth.access"`
}

var DBConfig DBEnv
var WebAPIConfig WebAPIEnv
var BotAPIConfig BotAPIEnv

func init() {
	if err := envconfig.Process("", &DBConfig); err != nil {
		log.Fatalf("Failed to process env var : %s", err)
		return
	}
	if err := envconfig.Process("", &WebAPIConfig); err != nil {
		log.Fatalf("Failed to process env var : %s", err)
		return
	}
	if err := envconfig.Process("", &BotAPIConfig); err != nil {
		log.Fatalf("Failed to process env var : %s", err)
		return
	}
}