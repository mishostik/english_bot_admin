package usecase

import (
	"context"
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/user"
	"english_bot_admin/pkg/secure"
	"fmt"
	"github.com/google/uuid"
	"time"
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

func (u *UserUsecase) AdminSignUp(context_ context.Context, params *models.AdminSignInParams) error {
	admin_ := &models.Admin{
		AdminId:      uuid.New(),
		Login:        params.Login,
		Password:     secure.HidePassword(params.Password),
		RegisteredAt: time.Now(),
	}

	err := u.repo.InsertAdmin(context_, admin_)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) AdminSignIn(context_ context.Context, params *models.AdminSignInParams) (bool, error) {
	params.Password = secure.HidePassword(params.Password)
	isExist, err := u.repo.AdminVerification(context_, params)
	if err != nil {
		return false, err
	}
	if !isExist {
		return false, fmt.Errorf("error: no admin")
	}
	return isExist, nil
}
