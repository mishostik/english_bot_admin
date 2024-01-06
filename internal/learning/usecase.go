package learning

import "context"

type Usecase interface {
	AddRule(ctx context.Context, rule *NewRuleParams) error
	GetAllRules(ctx context.Context) ([]Rule, error)
}
