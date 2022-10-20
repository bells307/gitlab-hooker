package usecase

import (
	"log"

	"github.com/bells307/gitlab-hooker/internal/model"
)

type MergeRequestUsecase interface {
	ProcessMergeRequest(mr *model.MergeRequest) error
}

type mergeRequestUsecase struct {
}

func NewMergeRequestUsecase() *mergeRequestUsecase {
	return &mergeRequestUsecase{}
}

func (u *mergeRequestUsecase) ProcessMergeRequest(mr *model.MergeRequest) error {
	log.Printf("mr: %v", mr)
	return nil
}
