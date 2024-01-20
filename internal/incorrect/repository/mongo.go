package repository

import (
	"context"
	"english_bot_admin/internal/incorrect"
	"english_bot_admin/internal/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type IncorrectRepository struct {
	incAnswers *mongo.Collection
}

func NewIncorrectRepository(incAnswers *mongo.Collection) incorrect.Repository {
	return &IncorrectRepository{
		incAnswers: incAnswers,
	}
}

func (r *IncorrectRepository) AddForNewTask(ctx context.Context, taskId uuid.UUID, a string, b string, c string) error {
	result := &models.IncorrectAnswers{
		TaskID: taskId,
		A:      a,
		B:      b,
		C:      c,
	}
	_, err := r.incAnswers.InsertOne(ctx, result)
	if err != nil {
		return fmt.Errorf("error while adding task: %w", err)
	}
	log.Println("incorrect answers for new task added")
	return nil
}

func (r *IncorrectRepository) UpdateForTask(ctx context.Context, taskId uuid.UUID, answers *models.IncorrectAnswers) error {
	filter := bson.M{"task_id": taskId}
	update := bson.M{
		"$set": bson.M{
			"a": answers.A,
			"b": answers.B,
			"c": answers.C,
		},
	}
	res, err := r.incAnswers.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		err := r.AddForNewTask(ctx, taskId, answers.A, answers.B, answers.C)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *IncorrectRepository) GetAnswersForTask(ctx context.Context, taskId uuid.UUID) (*models.IncorrectAnswers, error) {
	filter := bson.M{"task_id": taskId}
	var answers models.IncorrectAnswers
	err := r.incAnswers.FindOne(ctx, filter).Decode(&answers)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &models.IncorrectAnswers{
				TaskID: taskId,
				A:      "empty",
				B:      "empty",
				C:      "empty",
			}, nil
		}
		return nil, err
	}
	return &answers, nil
}
