package application

import (
	"fmt"

	"github.com/bells307/gitlab-hooker/internal/domain"
)

type hookService struct {
	senderService SenderService
}

// Сервис отправки сообщений
type SenderService interface {
	SendMessageToChats(string)
}

func NewHookService(senderService SenderService) *hookService {
	return &hookService{senderService}
}

func (s *hookService) ProcessMergeRequestHook(mr domain.MergeRequestHook) error {
	// TODO: templates
	if mr.ObjectAttributes.State == "opened" && mr.ObjectAttributes.Action == "open" {
		msg := fmt.Sprintf(
			"🔍 <b>%s</b> создал <b>Merge Request</b> \"%s\" на проекте <i>%s</i>:\n%s",
			mr.User.Name,
			mr.ObjectAttributes.Title,
			mr.Project.Name,
			mr.ObjectAttributes.URL,
		)
		s.senderService.SendMessageToChats(msg)
	} else if mr.ObjectAttributes.State == "merged" && mr.ObjectAttributes.Action == "merge" {
		msg := fmt.Sprintf(
			"✅ <b>%s</b> слил изменения по <b>Merge Request</b> \"%s\" на проекте <i>%s</i>:\n%s",
			mr.User.Name,
			mr.ObjectAttributes.Title,
			mr.Project.Name,
			mr.ObjectAttributes.URL,
		)
		s.senderService.SendMessageToChats(msg)
	}

	return nil
}
