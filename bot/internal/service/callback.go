package service

import (
	"fmt"
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

const offset = 10

type PagerData struct {
	Offset      int    `json:"offset"`
	ButtonType  string `json:"button_type"`
	CurrentPage int    `json:"current_page"`
	MaxPages    int
	Products    []models.Product
}

func (b *Bot) Pager(products []models.Product, callbackBody string) (string, tgbotapi.InlineKeyboardMarkup) {
	pagerData := PagerData{
		Offset:   offset,
		Products: products,
	}
	ParseCallbackPagerData(callbackBody, &pagerData)
	pagerData.MaxPages = len(pagerData.Products) / offset
	if pagerData.ButtonType == "next" {
		nextPage := pagerData.CurrentPage + 1
		if nextPage <= pagerData.MaxPages {
			pagerData.CurrentPage = nextPage
			return RenderProductListWithMarkup(&pagerData)
		}
	}
	if pagerData.ButtonType == "previous" {
		previousPage := pagerData.CurrentPage - 1
		if previousPage >= 0 {
			pagerData.CurrentPage = previousPage
			return RenderProductListWithMarkup(&pagerData)
		}
	}

	return RenderProductListWithMarkup(&pagerData)
}

func RenderProductListWithMarkup(pagerData *PagerData) (string, tgbotapi.InlineKeyboardMarkup) {
	markUp := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", fmt.Sprintf("pager_next:%d:%d", pagerData.CurrentPage, offset)),
		),
	)
	text := GenerateTextForBotMessage(pagerData.Products[pagerData.CurrentPage*offset : pagerData.CurrentPage*offset+offset])
	if pagerData.CurrentPage > 0 && pagerData.CurrentPage < pagerData.MaxPages {
		markUp = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Previous page", fmt.Sprintf("pager_previous:%d:%d", pagerData.CurrentPage, offset)),
				tgbotapi.NewInlineKeyboardButtonData("Next page", fmt.Sprintf("pager_next:%d:%d", pagerData.CurrentPage, offset)),
			),
		)
	}
	if pagerData.CurrentPage == pagerData.MaxPages {
		markUp = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Previous page", fmt.Sprintf("pager_previous:%d:%d", pagerData.CurrentPage, offset)),
			),
		)
	}
	return text, markUp
}

func ParseCallbackPagerData(callbackBody string, pagerData *PagerData) {
	parsedData := strings.Split(callbackBody, ":")
	pagerData.ButtonType = parsedData[0]
	pagerData.CurrentPage, _ = strconv.Atoi(parsedData[1])
}

func GenerateTextForBotMessage(products []models.Product) string {
	str := ""
	for _, value := range products {
		if (value != models.Product{}) {
			str += fmt.Sprintf("%v\n", value)
		}
	}
	return str
}
