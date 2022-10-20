package service

import (
	"log"

	"github.com/bells307/gitlab-hooker/internal/model"
)

type MergeRequestService interface {
	ProcessMergeRequest(mr *model.MergeRequest) error
}

type mergeRequestService struct {
	telegramService TelegramService
}

func NewMergeRequestService(telegramService TelegramService) *mergeRequestService {
	return &mergeRequestService{telegramService}
}

func (s *mergeRequestService) ProcessMergeRequest(mr *model.MergeRequest) error {
	log.Printf("mr: %v", mr)
	return nil
}
