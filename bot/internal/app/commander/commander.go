package commander

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	bot *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI,productService *product.Service)*Commander{
	return &Commander{
		bot: bot,
		productService: productService,
	}
}