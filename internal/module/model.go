package module

import (
	"github.com/google/uuid"
)

type Module struct {
	ModuleID uuid.UUID    `json:"module_id" bson:"module_id"`
	Title    string       `json:"title" bson:"title"`
	Level    string       `json:"level" bson:"level"`
	Task     *[]uuid.UUID `json:"task" bson:"task,omitempty"`
}

type NewModuleParams struct {
	Title string       `json:"title"`
	Level string       `json:"level" bson:"level"`
	Task  *[]uuid.UUID `json:"task" bson:"task,omitempty"`
}
