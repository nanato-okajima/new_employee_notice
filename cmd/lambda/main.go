package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"new_employee_notice/internal/batch"
	"new_employee_notice/internal/config"
	"new_employee_notice/internal/discord"
	"new_employee_notice/internal/redis"
	"new_employee_notice/internal/scraping"
)

var (
	envPath = "build/env/.local.env"
)

func HandleRequest(ctx context.Context) {
	if err := config.Setup(envPath); err != nil {
		log.Fatal(err)
	}

	n := redis.New()
	s := scraping.New()
	d, err := discord.New()
	if err != nil {
		log.Fatal(err)
	}

	b := batch.New(s, d, n)
	log.Fatal(b.Exec())
}

func main() {
	lambda.Start(HandleRequest)
}
