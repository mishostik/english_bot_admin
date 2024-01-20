package usecase

import (
	"context"
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/user"
)

type UserUsecase struct {
	repo user.Repository
}

func NewUserUsecase(repo user.Repository) user.Usecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) GetAll(context_ context.Context) ([]models.User, error) {
	users, err := u.repo.Select(context_)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}
