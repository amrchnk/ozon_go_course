package service

import (
	"encoding/json"
	"fmt"
	"github.com/amrchnk/ozon_go_course/bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	Bot             *tgbotapi.BotAPI
	ProductRepository repository.ProductRepository
}

type CommandData struct {
	Offset int `json:"offset"`
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
			parsedData:= CommandData{}
			json.Unmarshal([]byte(update.CallbackQuery.Data),&parsedData)
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Parsed: %v\n",parsedData))
			b.Bot.Send(msg)
			return
		}

		if update.Message == nil { // If we got a message
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

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.Bot.GetUpdatesChan(u)
}