package commander

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander)List(inputMessage *tgbotapi.Message) {
	out:=""
	products:=c.productService.List()
	for _,p:=range products{
		out+=p.Title+"\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Products list:\n"+out)

	serializedData,_:=json.Marshal(CommandData{
		Offset: 10,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)
	c.bot.Send(msg)
}
