package service

import (
	"fmt"

	"github.com/bells307/gitlab-hooker/internal/interfaces"
	"github.com/bells307/gitlab-hooker/internal/model"
)

type hookService struct {
	senderService interfaces.SenderService
}

func NewHookService(senderService interfaces.SenderService) *hookService {
	return &hookService{senderService}
}

func (s *hookService) ProcessMergeRequestHook(mr model.MergeRequestHook) error {
	var msg string

	if mr.State == "opened" && mr.Action == "open" {
		msg = fmt.Sprintf("<b>%s</b> создал <b>Merge Request</b> \"%s\" на проекте <i>%s</i>:\n%s", mr.Username, mr.Title, mr.Project, mr.Url)
	} else if mr.State == "merged" && mr.Action == "merge" {
		msg = fmt.Sprintf("<b>%s</b> слил изменения по <b>Merge Request</b> \"%s\" на проекте <i>%s</i>:\n%s", mr.Username, mr.Title, mr.Project, mr.Url)
	}

	if msg == "" {
		return nil
	}

	s.senderService.SendMessageToChats(msg)
	return nil
}
