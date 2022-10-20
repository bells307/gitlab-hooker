package service

import (
	tm "github.com/and3rson/telemux/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService interface{}

type telegramService struct {
	api *tgbotapi.BotAPI
}

type TelegramServiceConfig struct {
	Token string
	Debug bool
}

func NewTelegramService(cfg TelegramServiceConfig) (*telegramService, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		api.Debug = true
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	srv := telegramService{api}
	go srv.run()

	return &srv, nil
}

func (t *telegramService) run() {
	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updChan := t.api.GetUpdatesChan(updConfig)

	mux := tm.NewMux()
	for upd := range updChan {
		mux.Dispatch(t.api, upd)
	}
}
