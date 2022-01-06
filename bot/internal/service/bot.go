package service

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
	"github.com/amrchnk/ozon_go_course/bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	Bot             *tgbotapi.BotAPI
	ProductRepository repository.ProductRepository
}

func NewBot(bot *tgbotapi.BotAPI,productRepository repository.ProductRepository) *Bot {
	return &Bot{
		Bot: bot,
		ProductRepository: productRepository,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.Bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	b.handleUpdates(updates)
	return nil
}

func(b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel){
	defer func(){
		if panicValue:=recover();panicValue!=nil{
			log.Printf("recovered from panic: %v",panicValue)
		}
	}()

	for update := range updates {
		if update.CallbackQuery!=nil{
			b.HandleCallback(update.CallbackQuery)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message!=nil{
			if err:=b.handleCommand(update.Message);err!=nil{
				log.Println(err)
				return
			}
			continue
		}
	}
}

func (b *Bot)HandleCallback(query *tgbotapi.CallbackQuery){
	callback,err:=models.ParseCallback(query.Data)
	products, err := b.ProductRepository.GetProductList()
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback.go data `%s` - %v", query.Data, err)
		return
	}
	switch callback.CallbackName {
	case "pager":
		text,markup:=b.Pager(products,callback.CallbackData)
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID,query.Message.MessageID, text)
		msg.ReplyMarkup = &markup
		b.Bot.Send(msg)
		return
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.Bot.GetUpdatesChan(u)
}

