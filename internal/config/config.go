package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Conf Config

type Config struct {
	MyPageUrl     string `envconfig:"MYPAGE_URL"`
	LoginUrl      string `envconfig:"LOGIN_URL"`
	LoginId       string `envconfig:"LOGIN_ID"`
	LoginPassword string `envconfig:"LOGIN_PASSWORD"`
	BotToken      string `envconfig:"BOT_TOKEN"`
	ChannelId     string `envconfig:"CHANNEL_ID"`
}

func Setup(envPath string) error {
	if err := godotenv.Load(envPath); err != nil {
		return err
	}

	if err := envconfig.Process("", &Conf); err != nil {
		return err
	}
	return nil
}
