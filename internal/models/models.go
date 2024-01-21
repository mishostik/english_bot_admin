package models

import (
	"github.com/google/uuid"
	"time"
)

type IncorrectAnswers struct {
	TaskID uuid.UUID `bson:"task_id"`
	A      string    `bson:"a"`
	B      string    `bson:"b"`
	C      string    `bson:"c"`
}

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

type RulesResponseModel struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    []Rule `json:"data"`
}

type Module struct {
	ModuleID uuid.UUID `json:"module_id" bson:"module_id"`
	Title    string    `json:"title" bson:"title"`
	Level    string    `json:"level" bson:"level"`
	Task     []Task    `json:"task" bson:"task,omitempty"`
}

type NewModuleParams struct {
	Title string `json:"title"`
	Level string `json:"level" bson:"level"`
	//Task  *[]uuid.UUID `json:"task" bson:"task,omitempty"`
}

type Lvl struct {
	Level    string    `json:"level"`
	ModuleID uuid.UUID `json:"module_id" bson:"module_id"`
}

type TaskToModule struct {
	ModuleId uuid.UUID `json:"module_id" bson:"module_id"`
	//TaskId   uuid.UUID `json:"task_id" bson:"task_id"`
	Task *Task `json:"task" bson:"task"`
}

type AddTaskByLvlParams struct {
	TaskId   uuid.UUID `json:"task_id"`
	ModuleId uuid.UUID `json:"module_id"`
}

type Task struct {
	TaskID   uuid.UUID `bson:"task_id" json:"task_id,omitempty"`
	TypeID   uint8     `bson:"type_id" json:"type_id"`
	Level    string    `bson:"level" json:"level"`
	Question string    `bson:"question" json:"question"`
	Answer   string    `bson:"answer" json:"answer"`
}

type TaskWithAnswers struct {
	TaskID   uuid.UUID `bson:"task_id" json:"task_id,omitempty"`
	TypeID   uint8     `bson:"type_id" json:"type_id"`
	Level    string    `bson:"level" json:"level"`
	Question string    `bson:"question" json:"question"`
	Answer   string    `bson:"answer" json:"answer"`
	A        string    `bson:"a" json:"a"`
	B        string    `bson:"b" json:"b"`
	C        string    `bson:"c" json:"c"`
}

type TasksResponseModel struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    []Task `json:"data"`
}

type TasksByLvlResponseModel struct {
	Success bool       `json:"success"`
	Error   string     `json:"error"`
	Data    []ByModule `json:"data"`
}

type ResponseModel struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ModulesResponseModel struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Data    []Module `json:"data"`
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

type User struct {
	UserID       int       `json:"user_id" bson:"user_id"`
	Username     string    `json:"username" bson:"username"`
	RegisteredAt time.Time `json:"registered_at" bson:"registered_at"`
	LastActiveAt time.Time `json:"last_active_at" bson:"last_active_at"`
	Level        string    `json:"level" bson:"level"`
}

type AdminSignInParams struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}

type Admin struct {
	AdminId      uuid.UUID `bson:"admin_id"`
	Login        string    `bson:"login"`
	Password     string    `bson:"password"`
	RegisteredAt time.Time `bson:"registered_at"`
}

type UsersResponseModel struct {
	Data    []User `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
