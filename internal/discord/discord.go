package discord

import (
	"log"
	"new_employee_notice/internal/config"

	"github.com/bwmarrin/discordgo"
)

type BotCli struct {
	discord   *discordgo.Session
	channelID string
}

type DiscordBot interface {
	SendEmployeeInfo(msg string) error
}

func New() (*BotCli, error) {
	d, err := discordgo.New(config.Conf.BotToken)
	if err != nil {
		return nil, err
	}
	d.Token = config.Conf.BotToken

	if err := d.Open(); err != nil {
		return nil, err
	}

	return &BotCli{
		discord:   d,
		channelID: config.Conf.ChannelId,
	}, nil
}

func (b BotCli) SendEmployeeInfo(msg string) error {
	if _, err := b.discord.ChannelMessageSend(b.channelID, msg); err != nil {
		return err
	}
	log.Println(">>>" + msg)

	return nil
}
