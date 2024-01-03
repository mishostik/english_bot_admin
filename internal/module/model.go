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

type Lvl struct {
	Level    string    `json:"level"`
	ModuleID uuid.UUID `json:"module_id" bson:"module_id"`
}

type TaskToModule struct {
	ModuleId uuid.UUID `json:"module_id" bson:"module_id"`
	TaskId   uuid.UUID `json:"task_id" bson:"task_id"`
}
