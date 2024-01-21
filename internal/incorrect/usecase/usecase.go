package usecase

import (
	"context"
	"english_bot_admin/internal/incorrect"
	"english_bot_admin/internal/models"
	"github.com/google/uuid"
)

type IncorrectUsecase struct {
	repo incorrect.Repository
}

func NewIncorrectUsecase(repo incorrect.Repository) incorrect.Usecase {
	return &IncorrectUsecase{
		repo: repo,
	}
}

func (u *IncorrectUsecase) AddForTask(ctx context.Context, taskId uuid.UUID, a string, b string, c string) error {
	err := u.repo.InsertForNewTask(ctx, taskId, a, b, c)
	if err != nil {
		return err
	}
	return nil
}

func (u *IncorrectUsecase) UpdateForTask(ctx context.Context, taskId uuid.UUID, answers *models.IncorrectAnswers) error {
	err := u.repo.UpdateForTask(ctx, taskId, answers)
	if err != nil {
		return err
	}
	return nil
}

func (u *IncorrectUsecase) GetAnswersForTask(ctx context.Context, taskId uuid.UUID) (*models.IncorrectAnswers, error) {
	task, err := u.repo.SelectAnswersForTask(ctx, taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}
