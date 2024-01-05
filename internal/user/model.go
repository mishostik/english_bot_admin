package user

import (
	"time"
)

type User struct {
	UserID       int       `json:"user_id" bson:"user_id"`
	Username     string    `json:"username" bson:"username"`
	RegisteredAt time.Time `json:"registered_at" bson:"registered_at"`
	LastActiveAt time.Time `json:"last_active_at" bson:"last_active_at"`
	Level        string    `json:"level" bson:"level"`
}
