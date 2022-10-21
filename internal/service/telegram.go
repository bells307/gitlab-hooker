package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService interface {
	AddedToChat(*tgbotapi.Update)
	RemovedFromChat(*tgbotapi.Update)
}

type telegramService struct {
	api *tgbotapi.BotAPI
}

func NewTelegramService(api *tgbotapi.BotAPI) *telegramService {
	return &telegramService{api}
}

func (t *telegramService) AddedToChat(*tgbotapi.Update) {
	log.Print("i was added")
}

func (t *telegramService) RemovedFromChat(*tgbotapi.Update) {
	log.Print("i was removed")
}
