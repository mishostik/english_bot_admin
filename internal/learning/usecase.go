package learning

import (
	"context"
	"english_bot_admin/internal/models"
)

type Usecase interface {
	AddRule(ctx context.Context, rule *models.NewRuleParams) error
	GetAllRules(ctx context.Context) ([]models.Rule, error)
}
