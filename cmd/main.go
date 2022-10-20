package main

import (
	controller "github.com/bells307/gitlab-hooker/internal/controller"
	"github.com/bells307/gitlab-hooker/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	telegramService, err := service.NewTelegramService(service.TelegramServiceConfig{
		Token: "123",
		Debug: true,
	})

	if err != nil {
		panic(err)
	}

	mergeRequestService := service.NewMergeRequestService(telegramService)

	router := gin.Default()
	h := controller.NewHookHandler(mergeRequestService)
	h.Register(router)

	router.Run("0.0.0.0:8888")
}
