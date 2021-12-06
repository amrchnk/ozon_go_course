package commander

import (
	"encoding/json"
	"fmt"
	"github.com/amrchnk/ozon_go_course/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander struct {
	bot *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Offset int `json:"offset"`
}

func NewCommander(bot *tgbotapi.BotAPI,productService *product.Service)*Commander{
	return &Commander{
		bot: bot,
		productService: productService,
	}
}

func (c *Commander)Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}

func (c *Commander)HandleUpdate(update tgbotapi.Update){
	defer func(){
		if panicValue:=recover();panicValue!=nil{
			fmt.Printf("recovered from panic: %v",panicValue)
		}
	}()

	if update.CallbackQuery!=nil{
		parsedData:=CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data),&parsedData)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Parsed: %v\n",parsedData))
		c.bot.Send(msg)
		return
	}

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