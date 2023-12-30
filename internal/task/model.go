package task

import (
	"github.com/google/uuid"
)

type Task struct {
	TaskID   uuid.UUID `bson:"task_id"`
	TypeID   uint8     `bson:"type_id" json:"type_id"`
	Level    string    `bson:"level" json:"level"`
	Question string    `bson:"question" json:"question"`
	Answer   string    `bson:"answer" json:"answer"`
}

type ToModule struct {
	ModuleID uuid.UUID `json:"module_id"`
	TaskID   uuid.UUID `json:"task_id"`
}
