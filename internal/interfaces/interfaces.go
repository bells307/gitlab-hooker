package interfaces

import (
	tm "github.com/and3rson/telemux/v2"
	"github.com/bells307/gitlab-hooker/internal/model"
)

// Сервис обработки хуков от гитлаба
type HookService interface {
	ProcessMergeRequestHook(model.MergeRequestHook) error
}

// Сервис отправки сообщений
type SenderService interface {
	SendMessageToChats(string)
}

// Сервис обработки Update'ов телеграма
type UpdateService interface {
	AddedToChat(*tm.Update)
	RemovedFromChat(*tm.Update)
}
