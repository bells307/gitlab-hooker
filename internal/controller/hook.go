package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bells307/gitlab-hooker/internal/model"
	"github.com/bells307/gitlab-hooker/internal/usecase"
	"github.com/gin-gonic/gin"
)

type hookHandler struct {
	mergeRequestUsecase usecase.MergeRequestUsecase
}

func NewHookHandler(mergeRequestUsecase usecase.MergeRequestUsecase) *hookHandler {
	return &hookHandler{mergeRequestUsecase}
}

func (h *hookHandler) Register(router *gin.Engine) {
	router.POST("/api/hook", h.processHook)
}

func (h *hookHandler) processHook(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var mr model.MergeRequest
	err = json.Unmarshal(jsonData, &mr)
	if err != nil {
		panic(err)
	}

	h.mergeRequestUsecase.ProcessMergeRequest(&mr)
}
