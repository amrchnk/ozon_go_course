package main

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/service/product"
	"github.com/amrchnk/ozon_go_course/bot/internal/app/commander"
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

	productService:=product.NewService()
	commander:=commander.NewCommander(bot,productService)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		command := update.Message.Command()
		switch command {
		case "help":
			commander.Help(update.Message)
			continue
		case "list":
			commander.List(update.Message)
			continue
		default:
			commander.Default(update.Message)
		}
	}
}

