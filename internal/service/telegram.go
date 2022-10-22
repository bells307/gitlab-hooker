package service

import (
	"log"

	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// Сервис для работы с телеграм-ботом
type telegramService struct {
	api   *tgbotapi.BotAPI
	chats []int64
}

func NewTelegramService(api *tgbotapi.BotAPI, chats []int64) *telegramService {
	return &telegramService{api, chats}
}

func (t *telegramService) AddedToChat(u *tm.Update) {
	chat_id := u.Update.MyChatMember.Chat.ID
	title := u.Update.MyChatMember.Chat.Title
	log.Printf("bot added to chat \"%s\", chat id: %d", title, chat_id)
	t.chats = append(t.chats, chat_id)
	// Обновляем конфигурацию
	viper.Set("Telegram.Chats", t.chats)
	viper.WriteConfig()
}

func (t *telegramService) RemovedFromChat(u *tm.Update) {
	chat_id := u.Update.MyChatMember.Chat.ID
	title := u.Update.MyChatMember.Chat.Title
	log.Printf("bot removed from chat \"%s\", chat id: %d", title, chat_id)

	idx := -1
	for i, c := range t.chats {
		if c == chat_id {
			idx = i
			break
		}
	}

	if idx != -1 {
		t.chats = append(t.chats[:idx], t.chats[idx+1:]...)
		// Обновляем конфигурацию
		viper.Set("Telegram.Chats", t.chats)
		viper.WriteConfig()
	}
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
