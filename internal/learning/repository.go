package learning

import "context"

type Repository interface {
	InsertRule(ctx context.Context, rule *Rule) error
	SelectRules(ctx context.Context) ([]Rule, error)
}
