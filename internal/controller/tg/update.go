package tg

import (
	tm "github.com/and3rson/telemux/v2"
	"github.com/bells307/gitlab-hooker/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type updateHandler struct {
	updateService service.TelegramService
}

func NewUpdateHandler(updateService service.TelegramService) *updateHandler {
	return &updateHandler{updateService}
}

func (h *updateHandler) Register(api *tgbotapi.BotAPI, mux *tm.Mux) {
	mux.AddHandler(tm.NewHandler(
		func(u *tm.Update) bool {
			if msg := u.EffectiveMessage(); msg != nil {
				if members := msg.NewChatMembers; members != nil {
					for _, m := range members {
						if m.ID == api.Self.ID {
							return true
						}
					}
				}
			}

			return false
		},
		func(u *tm.Update) {
			h.updateService.AddedToChat(&u.Update)
		},
	)).AddHandler(tm.NewCommandHandler(
		"help",
		nil,
		func(u *tm.Update) {
			api.Send(tgbotapi.NewMessage(u.Message.Chat.ID, "help"))
		},
	))
}
