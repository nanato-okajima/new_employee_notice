package batch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v9"

	"new_employee_notice/internal/discord"
	r "new_employee_notice/internal/redis"
	"new_employee_notice/internal/scraping"
)

const (
	newCommerMessage    = "新しく入った人"
	infoMessageTemplate = "%s(%s) \n 生年月日:%s\n 研修開始日:%s\n"
)

type batchCli struct {
	scraping scraping.Scraping
	bot      discord.DiscordBot
	noSQL    r.NoSQL
}

func New(s scraping.ScrapingCli, b *discord.BotCli, n r.NoSQL) batchCli {
	return batchCli{
		scraping: &s,
		bot:      b,
		noSQL:    n,
	}
}

func (b batchCli) Exec() error {
	ctx := context.Background()
	emps, err := b.scraping.FetchEmployeeData()
	if err != nil {
		return err
	}

	firstTime := true
	for _, emp := range emps {
		if _, err := b.noSQL.Get(ctx, emp.ID); err == nil {
			continue
		} else if err != nil && err != redis.Nil {
			return err
		}

		serialize, err := json.Marshal(&emp)
		if err != nil {
			return err
		}

		if err := b.noSQL.SetNX(ctx, emp.ID, serialize, 0); err != nil {
			return err
		}

		if firstTime {
			firstTime = false
			if err := b.bot.SendEmployeeInfo(newCommerMessage); err != nil {
				return err
			}
		}

		msg := fmt.Sprintf(infoMessageTemplate, emp.Name, emp.Furigana, emp.BirthDay, emp.TrainingStartDate)
		if err := b.bot.SendEmployeeInfo(msg); err != nil {
			return err
		}
	}
	return nil
}
