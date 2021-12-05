package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander)List(inputMessage *tgbotapi.Message) {
	out:=""
	products:=c.productService.List()
	for _,p:=range products{
		out+=p.Title+"\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Products list:\n"+out)
	c.bot.Send(msg)
}
