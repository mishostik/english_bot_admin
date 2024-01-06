package learning

import (
	"github.com/google/uuid"
)

type Rule struct {
	RuleID   uuid.UUID `json:"rule_id" bson:"rule_id"`
	Info     string    `json:"info" bson:"info"`
	Image    []byte    `json:"image" bson:"image"`
	ModuleID uuid.UUID `json:"module_id" bson:"module_id,omitempty"`
	Topic    string    `json:"topic" bson:"topic"`
}

type NewRuleParams struct {
	Info     string    `json:"info" bson:"info"`
	Image    []byte    `json:"image" bson:"image"`
	ModuleID uuid.UUID `json:"module_id" bson:"module_id,omitempty"`
	Topic    string    `json:"topic" bson:"topic"`
}
