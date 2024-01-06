package usecase

import (
	"context"
	"english_bot_admin/internal/learning"
	"fmt"
	"github.com/google/uuid"
)

type LearnUseCase struct {
	repo learning.Repository
}

func NewLearnUsecase(learnRepo learning.Repository) learning.Usecase {
	return &LearnUseCase{repo: learnRepo}
}

func (u *LearnUseCase) AddRule(ctx context.Context, rule *learning.NewRuleParams) error {
	temp := &learning.Rule{
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

func (u *LearnUseCase) GetAllRules(ctx context.Context) ([]learning.Rule, error) {
	var (
		rules []learning.Rule
		err   error
	)

	rules, err = u.repo.SelectRules(ctx)
	if err != nil {
		return rules, fmt.Errorf("error getting rules: %v", err.Error())
	}
	return rules, nil
}
