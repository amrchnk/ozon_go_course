package main

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/app/commander"
	"github.com/amrchnk/ozon_go_course/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()
	commander := commander.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
