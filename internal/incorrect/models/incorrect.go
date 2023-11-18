package models

import (
	"github.com/google/uuid"
)

type IncorrectAnswers struct {
	TaskID uuid.UUID `bson:"task_id"`
	A      string    `bson:"a"`
	B      string    `bson:"b"`
	C      string    `bson:"c"`
}
