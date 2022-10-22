package http

import (
	"encoding/json"
	"io"

	"github.com/bells307/gitlab-hooker/internal/interfaces"
	"github.com/bells307/gitlab-hooker/internal/model"
	"github.com/gin-gonic/gin"
)

type hookHandler struct {
	hookService interfaces.HookService
}

func NewHookHandler(hookService interfaces.HookService) *hookHandler {
	return &hookHandler{hookService}
}

func (h *hookHandler) Register(router *gin.Engine) {
	router.POST("/api/hook", h.processHook)
}

func (h *hookHandler) processHook(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]any
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		panic(err)
	}

	objectAttributes := result["object_attributes"].(map[string]any)
	user := result["user"].(map[string]any)
	project := result["project"].(map[string]any)
	objectKind := result["object_kind"].(string)
	if objectKind == "merge_request" {
		hook := model.MergeRequestHook{
			Title:     objectAttributes["title"].(string),
			State:     objectAttributes["state"].(string),
			Action:    objectAttributes["action"].(string),
			Username:  user["name"].(string),
			Url:       objectAttributes["url"].(string),
			Project:   project["name"].(string),
			Assignees: []string{},
		}
		err = h.hookService.ProcessMergeRequestHook(hook)
		if err != nil {
			panic(err)
		}
	} else {
		panic("unknown hook")
	}
}
