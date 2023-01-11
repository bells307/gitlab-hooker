package http

import (
	"errors"

	"github.com/bells307/gitlab-hooker/internal/domain"
	"github.com/bells307/gitlab-hooker/pkg/gin/err_resp"
	"github.com/gin-gonic/gin"
)

const GitlabEventHeaderName string = "X-Gitlab-Event"

// Строковые значения, передаваемые в хедере
const (
	MergeRequestEventHeader string = "Merge Request Hook"
)

type hookHandler struct {
	hookService HookService
}

// Сервис обработки хуков от гитлаба
type HookService interface {
	ProcessMergeRequestHook(domain.MergeRequestHook) error
}

func NewHookHandler(hookService HookService) *hookHandler {
	return &hookHandler{hookService}
}

func (h *hookHandler) Register(router *gin.Engine) {
	router.POST("/api/hook", h.processHook)
}

func (h *hookHandler) processHook(c *gin.Context) {
	header, ok := c.Request.Header[GitlabEventHeaderName]
	if !ok {
		err_resp.NewErrorResponse(c, errors.New("gitlab event header not provided"))
		return
	}

	switch header[0] {
	case MergeRequestEventHeader:
		var hook domain.MergeRequestHook
		if err := c.Bind(&hook); err != nil {
			return
		}

		if err := h.hookService.ProcessMergeRequestHook(hook); err != nil {
			err_resp.NewErrorResponse(c, err)
			return
		}
	default:
		err_resp.NewErrorResponse(c, errors.New("unknown event type header"))
		return
	}
}
