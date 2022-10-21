package main

import (
	tm "github.com/and3rson/telemux/v2"
	http "github.com/bells307/gitlab-hooker/internal/controller/http"
	tg "github.com/bells307/gitlab-hooker/internal/controller/tg"
	"github.com/bells307/gitlab-hooker/internal/service"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramBotApiConfig struct {
	Token string
	Debug bool
}

func main() {
	cfg := telegramBotApiConfig{
		Token: "5703775990:AAGhq55pBLEO4qzciJYqpivBPcCZDHEUwT4",
		Debug: true,
	}

	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		panic(err)
	}

	if cfg.Debug {
		api.Debug = true
	}

	telegramService := service.NewTelegramService(api)
	updateHandler := tg.NewUpdateHandler(telegramService)
	mux := tm.NewMux()
	updateHandler.Register(api, mux)

	mergeRequestService := service.NewMergeRequestService(telegramService)

	router := gin.Default()
	hookHandler := http.NewHookHandler(mergeRequestService)
	hookHandler.Register(router)

	go runTelegramBot(api, mux)
	router.Run("0.0.0.0:8888")
}

func runTelegramBot(api *tgbotapi.BotAPI, mux *tm.Mux) {
	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updChan := api.GetUpdatesChan(updConfig)

	for upd := range updChan {
		// log.Printf("new upd: %v", upd)
		// msg := tgbotapi.NewMessage(-1001874758944, "i got update")
		// api.Send(msg)
		mux.Dispatch(api, upd)
	}
}
