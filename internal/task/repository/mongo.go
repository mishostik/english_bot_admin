package task

import (
	"context"
	model "english_bot_admin/internal/task"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	taskCollection *mongo.Collection
	typeCollection *mongo.Collection
}

func NewMongoTaskRepository(taskCollection *mongo.Collection, typeCollection *mongo.Collection) model.Repository {
	return &MongoTaskRepository{
		taskCollection: taskCollection,
		typeCollection: typeCollection,
	}
}

func (r *MongoTaskRepository) UpdateTask(taskID int, task *model.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *MongoTaskRepository) GetTasks(ctx context.Context) ([]model.Task, error) {
	filter := bson.M{}
	var tasks []model.Task
	cursor, err := r.taskCollection.Find(ctx, filter)
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

func (r *MongoTaskRepository) GetTaskByID(ctx context.Context, taskID uuid.UUID) (*model.Task, error) {
	filter := bson.M{"task_id": taskID}
	var task model.Task
	err := r.taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("error no documents")
			return nil, nil
		}
		return nil, err
	}
	return &task, err
}

func (r *MongoTaskRepository) InsertTask(ctx context.Context, task *model.Task) (uuid.UUID, error) {
	_, err := r.taskCollection.InsertOne(ctx, task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error while adding task: %w", err)
	}
	log.Printf("Task added with ID: %s\n", task.TaskID)
	return task.TaskID, nil
}

func (r *MongoTaskRepository) UpdateTaskInfoByUUID(ctx context.Context, task *model.Task) error {
	filter := bson.M{"task_id": task.TaskID}
	update := bson.M{
		"$set": bson.M{
			"type_id":  task.TypeID,
			"level":    task.Level,
			"question": task.Question,
			"answer":   task.Answer,
			"task_id":  task.TaskID,
		},
	}
	_, err := r.taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("updating task in mongo db error")
		return err
	}
	return nil
}

func (r *MongoTaskRepository) UpdateTaskUUID(ctx context.Context, task *model.Task) error {
	filter := bson.M{"question": task.Question}
	update := bson.M{
		"$set": bson.M{
			"type_id":  task.TypeID,
			"level":    task.Level,
			"question": task.Question,
			"answer":   task.Answer,
			"task_id":  task.TaskID,
		},
	}

	_, err := r.taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoTaskRepository) DeleteTask(ctx context.Context, taskID int) error {
	filter := bson.M{"_id": taskID}
	_, err := r.taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoTaskRepository) GetTasksWithoutUUID(ctx context.Context) ([]model.Task, error) {
	filter := bson.M{"task_id": nil}
	var tasks []model.Task

	cursor, err := r.taskCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = cursor.Close(ctx)
	}()

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *MongoTaskRepository) GetTasksByLvl(ctx context.Context, lvl string) ([]model.Task, error) {
	log.Println("level is:", lvl)
	filter := bson.M{"level": lvl}
	var tasks []model.Task

	cursor, err := r.taskCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = cursor.Close(ctx)
	}()

	if err = cursor.All(ctx, &tasks); err != nil {
		log.Println("err or len:", err, tasks)
		return nil, err
	}

	return tasks, nil
}
