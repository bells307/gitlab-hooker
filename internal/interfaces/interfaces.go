package interfaces

import (
	"github.com/bells307/gitlab-hooker/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	AddedToChat(*tgbotapi.Update)
	RemovedFromChat(*tgbotapi.Update)
}
