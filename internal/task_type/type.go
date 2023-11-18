package task_type

type TaskType struct {
	TypeID int    `bson:"type_id"`
	Type   string `bson:"type"`
}
