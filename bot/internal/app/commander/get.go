package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander)Get(inputMessage *tgbotapi.Message) {
	args:=inputMessage.CommandArguments()

	id,err:=strconv.Atoi(args)
	if err!=nil{
		log.Println("Wrong args ",args)
		return
	}

	product,err:=c.productService.Get(id)
	if err!=nil{
		log.Printf("Fail to get product with id %d: %v",id,err)
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
