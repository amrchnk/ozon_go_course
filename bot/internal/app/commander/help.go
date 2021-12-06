package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander)Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n/list - list products\n/get - some")
	c.bot.Send(msg)
}