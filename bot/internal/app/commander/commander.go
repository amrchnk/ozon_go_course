package commander

import (
	"fmt"
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

func (c *Commander)HandleUpdate(update tgbotapi.Update){
	defer func(){
		if panicValue:=recover();panicValue!=nil{
			fmt.Printf("recovered from panic: %v",panicValue)
		}
	}()

	if update.Message == nil { // If we got a message
		return
	}

	command := update.Message.Command()
	switch command {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}