package main

import (
	controller "github.com/bells307/gitlab-hooker/internal/controller"
	"github.com/bells307/gitlab-hooker/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	mergeRequestUsecase := usecase.NewMergeRequestUsecase()

	router := gin.Default()
	h := controller.NewHookHandler(mergeRequestUsecase)
	h.Register(router)

	router.Run("0.0.0.0:8888")
}
