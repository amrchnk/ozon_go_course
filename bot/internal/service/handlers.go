package service

import (
	"encoding/json"
	"fmt"
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

const (
	commandHelp = "help"
	commandList = "list"
	commandGet  = "get"
	commandCreate  = "create"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandHelp:
		b.Help(message)
	case commandList:
		b.List(message)
	case commandGet:
		b.Get(message)
	case commandCreate:
		b.Create(message)
	default:
		b.Default(message)
	}
	return nil
}

func (b *Bot) Default(inputMessage *tgbotapi.Message) error {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	_, err := b.Bot.Send(msg)

	return err
}

func (b *Bot) Help(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n/list - list products\n/get - some")
	_, err := b.Bot.Send(msg)

	return err
}

func (b *Bot) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args ", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Некорректный id")
		b.Bot.Send(msg)
	}

	product, err := b.ProductRepository.GetProductById(int(id))
	if err != nil {
		log.Printf("Fail to get product with id %d: %v", id, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Ошибка получения товара с id %d: %v", id, err))
		b.Bot.Send(msg)
	}

	if (product==models.Product{}){
		log.Printf("Товара с id=%v не существует", id)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Товара с id=%v не существует", id))
		b.Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprint(product))
	b.Bot.Send(msg)
	return
}

func (b *Bot) Create(inputMessage *tgbotapi.Message){
	product:=models.Product{}
	json.Unmarshal([]byte(inputMessage.CommandArguments()),&product)

	if err:=(product==models.Product{});err{
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Ошибка при чтении полей: невалидные данные")
		b.Bot.Send(msg)
		return
	}
	err := b.ProductRepository.CreateProduct(&product)
	if err!=nil{
		log.Println(err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Ошибка: "+err.Error())
		b.Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Товар успешно создан")
	b.Bot.Send(msg)
	return
}

func (b *Bot) List(inputMessage *tgbotapi.Message) {
	products,err := b.ProductRepository.GetProductList()
	str:=""
	for _,value :=range products{
		str+=fmt.Sprintf("%v\n",value)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, str)
	b.Bot.Send(msg)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
