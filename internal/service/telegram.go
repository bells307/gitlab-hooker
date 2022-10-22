package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Сервис для работы с телеграм-ботом
type telegramService struct {
	api   *tgbotapi.BotAPI
	chats []int64
}

func NewTelegramService(api *tgbotapi.BotAPI, chats []int64) *telegramService {
	return &telegramService{api, chats}
}

func (t *telegramService) AddedToChat(*tgbotapi.Update) {
	log.Print("i was added")
}

func (t *telegramService) RemovedFromChat(*tgbotapi.Update) {
	log.Print("i was removed")
}

func (t *telegramService) SendMessageToChats(msg string) {
	for _, c := range t.chats {
		msgConfig := tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:           c,
				ReplyToMessageID: 0,
			},
			Text:                  msg,
			DisableWebPagePreview: false,
			ParseMode:             "HTML",
		}
		t.api.Send(msgConfig)
	}
}
