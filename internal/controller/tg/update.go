package tg

import (
	tm "github.com/and3rson/telemux/v2"
	"github.com/bells307/gitlab-hooker/internal/interfaces"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type updateHandler struct {
	updateService interfaces.UpdateService
}

func NewUpdateHandler(updateService interfaces.UpdateService) *updateHandler {
	return &updateHandler{updateService}
}

func (h *updateHandler) Register(api *tgbotapi.BotAPI, mux *tm.Mux) {
	mux.
		// AddHandler(tm.NewHandler(
		// 	func (u *tm.Update) bool {
		// 		if message := u.EffectiveMessage(); message != nil {
		// 			if message.NewChatMembers != nil && len(message.NewChatMembers) > 0 {

		// 			}
		// 		}
		// 		return false
		// 	},
		// 	h.updateService.AddedToChat,
		// ))
		AddHandler(tm.NewHandler(
			func(u *tm.Update) bool {
				if mcm := u.MyChatMember; mcm != nil {
					ncm := mcm.NewChatMember
					if ncm.User != nil {
						if ncm.User.ID == api.Self.ID {
							if (ncm.Status != "left") && (ncm.Status != "kicked") {
								return true
							}
						}
					}
				}

				return false
			},
			func(u *tm.Update) {
				h.updateService.AddedToChat(u)
			},
		)).
		AddHandler(tm.NewHandler(
			func(u *tm.Update) bool {
				if mcm := u.MyChatMember; mcm != nil {
					ncm := mcm.NewChatMember
					if ncm.User != nil {
						if ncm.User.ID == api.Self.ID {
							if (ncm.Status == "left") || (ncm.Status == "kicked") {
								return true
							}
						}
					}
				}

				return false
			},
			func(u *tm.Update) {
				h.updateService.RemovedFromChat(u)
			},
		))
}
