package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bells307/gitlab-hooker/internal/model"
	"github.com/bells307/gitlab-hooker/internal/service"
	"github.com/gin-gonic/gin"
)

type hookHandler struct {
	mergeRequestService service.MergeRequestService
}

func NewHookHandler(mergeRequestService service.MergeRequestService) *hookHandler {
	return &hookHandler{mergeRequestService}
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

	h.mergeRequestService.ProcessMergeRequest(&mr)
}
