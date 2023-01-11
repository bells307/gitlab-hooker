package sender

import (
	"log"

	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// Сервис для работы с телеграм-ботом
type telegramBot struct {
	api   *tgbotapi.BotAPI
	chats []int64
}

func NewTelegramBot(api *tgbotapi.BotAPI, chats []int64) *telegramBot {
	return &telegramBot{api, chats}
}

func (t *telegramBot) AddedToChat(u *tm.Update) {
	chatID := u.Update.MyChatMember.Chat.ID
	title := u.Update.MyChatMember.Chat.Title
	log.Printf("bot added to chat \"%s\", chat id: %d", title, chatID)
	t.chats = append(t.chats, chatID)
	// Обновляем конфигурацию
	viper.Set("Telegram.Chats", t.chats)
	viper.WriteConfig()
}

func (t *telegramBot) RemovedFromChat(u *tm.Update) {
	chatID := u.Update.MyChatMember.Chat.ID
	title := u.Update.MyChatMember.Chat.Title
	log.Printf("bot removed from chat \"%s\", chat id: %d", title, chatID)

	idx := -1
	for i, c := range t.chats {
		if c == chatID {
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

func (t *telegramBot) SendMessageToChats(msg string) error {
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
		if _, err := t.api.Send(msgConfig); err != nil {
			return err
		}

	}
	return nil
}
