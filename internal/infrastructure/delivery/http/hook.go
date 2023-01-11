package http

import (
	"errors"

	"github.com/bells307/gitlab-hooker/internal/domain/merge_request"
	"github.com/bells307/gitlab-hooker/internal/domain/pipeline"

	"github.com/bells307/gitlab-hooker/pkg/gin/err_resp"
	"github.com/gin-gonic/gin"
)

const GitlabEventHeaderName string = "X-Gitlab-Event"

// Строковые значения, передаваемые в хедере
const (
	MergeRequestEventHeader string = "Merge Request Hook"
	PipelineEventHeader     string = "Pipeline Hook"
)

type hookHandler struct {
	hookService HookService
}

// Сервис обработки хуков от гитлаба
type HookService interface {
	ProcessMergeRequestHook(merge_request.MergeRequest) error
	ProcessPipelineHook(pipeline.Pipeline) error
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
		var input MergeRequestHookInput
		if err := c.Bind(&input); err != nil {
			return
		}

		if err := h.hookService.ProcessMergeRequestHook(input.ToDomain()); err != nil {
			err_resp.NewErrorResponse(c, err)
			return
		}
	case PipelineEventHeader:
		var input PipelineHookInput
		if err := c.Bind(&input); err != nil {
			return
		}

		if err := h.hookService.ProcessPipelineHook(input.ToDomain()); err != nil {
			err_resp.NewErrorResponse(c, err)
			return
		}
	default:
		err_resp.NewErrorResponse(c, errors.New("unknown event type header"))
		return
	}
}
