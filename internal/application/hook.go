package application

import (
	"fmt"
	"github.com/bells307/gitlab-hooker/internal/domain/merge_request"
	"github.com/bells307/gitlab-hooker/internal/domain/pipeline"
	"github.com/bells307/gitlab-hooker/internal/infrastructure/sender"
)

// Сервис обработки хуков гитлаба
type hookService struct {
	sender sender.Sender
}

func NewHookService(sender sender.Sender) *hookService {
	return &hookService{sender}
}

func (s *hookService) ProcessMergeRequestHook(hook merge_request.MergeRequest) error {
	// TODO: templates
	if hook.State == merge_request.StateOpened && hook.Action == merge_request.ActionOpen {
		msg := fmt.Sprintf(
			"🔍 <b>%s</b> создал <b>Merge Request</b> \"%s\" на проекте <i>%s</i>:\n%s",
			hook.Username,
			hook.Title,
			hook.Project,
			hook.URL,
		)

		if err := s.sender.SendMessageToChats(msg); err != nil {
			return err
		}
	} else if hook.State == merge_request.StateMerged && hook.Action == merge_request.ActionMerge {
		msg := fmt.Sprintf(
			"✅ <b>%s</b> слил изменения по <b>Merge Request</b> \"%s\" на проекте <i>%s</i>:\n%s",
			hook.Username,
			hook.Title,
			hook.Project,
			hook.URL,
		)

		if err := s.sender.SendMessageToChats(msg); err != nil {
			return err
		}
	}

	return nil
}

func (s *hookService) ProcessPipelineHook(hook pipeline.Pipeline) error {
	if hook.Status == pipeline.Success {
		msg := fmt.Sprintf(
			"🔨 Успешно завершен <b>Pipeline</b> на проекте <i>%s</i> (ветка <i>%s</i>)",
			hook.Project,
			hook.Branch,
		)

		if err := s.sender.SendMessageToChats(msg); err != nil {
			return err
		}
	} else if hook.Status == pipeline.Failed {
		msg := fmt.Sprintf(
			"🧨 <b>Pipeline</b> завершился с ошибкой на проекте <i>%s</i> (ветка <i>%s</i>)",
			hook.Project,
			hook.Branch,
		)

		if err := s.sender.SendMessageToChats(msg); err != nil {
			return err
		}
	}

	return nil
}
