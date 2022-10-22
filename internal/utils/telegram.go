package utils

import (
	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotConfig struct {
	Token string  `mapstructure:"token"`
	Debug bool    `mapstructure:"debug"`
	Chats []int64 `mapstructure:"chats"`
}

func RunTelegramBot(api *tgbotapi.BotAPI, mux *tm.Mux) {
	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updChan := api.GetUpdatesChan(updConfig)

	for upd := range updChan {
		mux.Dispatch(api, upd)
	}
}
