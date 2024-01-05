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

// todo: wtf

type ToModule struct {
	ModuleID uuid.UUID `json:"module_id"`
	TaskID   uuid.UUID `json:"task_id"`
}

type ByModule struct {
	ModuleID uuid.UUID `bson:"module_id" json:"module_id"`
	TaskID   uuid.UUID `bson:"task_id"`
	Question string    `bson:"question" json:"question"`
	TypeID   uint8     `bson:"type_id" json:"type_id"`
}

type ByLvl struct {
	Level    string    `json:"level"`
	ModuleID uuid.UUID `json:"module_id" bson:"module_id"`
}
