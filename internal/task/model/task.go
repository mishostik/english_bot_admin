package task

type Task struct {
	TypeID   uint8  `bson:"type_id"`
	Level    string `bson:"level"`
	Question string `bson:"question"`
	Answer   string `bson:"answer"`
}
