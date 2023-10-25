package task

import (
	"context"
	model "english_bot_admin/internal/task/model"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	taskCollection *mongo.Collection
	typeCollection *mongo.Collection
}

func NewMongoTaskRepository(taskCollection *mongo.Collection, typeCollection *mongo.Collection) *MongoTaskRepository {
	return &MongoTaskRepository{
		taskCollection: taskCollection,
		typeCollection: typeCollection,
	}
}

func (mr *MongoTaskRepository) GetTasks(ctx context.Context) ([]model.Task, error) {
	filter := bson.M{}
	var tasks []model.Task
	cursor, err := mr.taskCollection.Find(ctx, filter)
	if err != nil {
		log.Println("error while getting tasks:", err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("")
		}
	}(cursor, ctx)
	if err := cursor.All(ctx, &tasks); err != nil {
		log.Println("error while decoding tasks:", err)
		return nil, err
	}
	return tasks, nil
}

func (mr *MongoTaskRepository) GetTaskByID(ctx context.Context, taskID int) (*model.Task, error) {
	filter := bson.M{"task_id": taskID}
	var task model.Task
	err := mr.taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("error no documents")
			return nil, nil
		}
		return nil, err
	}
	return &task, err
}

func (mr *MongoTaskRepository) NewTask(ctx context.Context, task *model.Task) error {
	_, err := mr.taskCollection.InsertOne(ctx, task)
	if err != nil {
		return fmt.Errorf("error while adding task: %w", err)
	}
	log.Println("task added")
	return nil
}

func (mr *MongoTaskRepository) UpdateTask(ctx context.Context, taskID int, task *model.Task) error {
	filter := bson.M{"task_id": taskID}
	update := bson.M{
		"$set": bson.M{
			"type_id":  task.TypeID,
			"level":    task.Level,
			"question": task.Question,
			"answer":   task.Answer,
		},
	}
	_, err := mr.taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("error while updating task:", err)
		return err
	}
	return nil
}

func (mr *MongoTaskRepository) DeleteTask(ctx context.Context, taskID int) error {
	filter := bson.M{"_id": taskID}
	_, err := mr.taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
