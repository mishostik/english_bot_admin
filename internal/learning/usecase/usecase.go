package usecase

import (
	"context"
	"english_bot_admin/internal/learning"
	"english_bot_admin/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type LearnUseCase struct {
	repo learning.Repository
}

func NewLearnUsecase(learnRepo learning.Repository) learning.Usecase {
	return &LearnUseCase{repo: learnRepo}
}

func (u *LearnUseCase) AddRule(ctx context.Context, rule *models.NewRuleParams) error {
	temp := &models.Rule{
		RuleID:   uuid.New(),
		ModuleID: rule.ModuleID,
		Info:     rule.Info,
		Image:    rule.Image,
		Topic:    rule.Topic,
	}
	err := u.repo.InsertRule(ctx, temp)
	if err != nil {
		return err
	}
	return nil
}

func (u *LearnUseCase) GetAllRules(ctx context.Context) ([]models.Rule, error) {
	var (
		rules []models.Rule
		err   error
	)

	rules, err = u.repo.SelectRules(ctx)
	if err != nil {
		return rules, fmt.Errorf("error getting rules: %v", err.Error())
	}
	return rules, nil
}
