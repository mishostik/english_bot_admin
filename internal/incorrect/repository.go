package incorrect

import (
	"context"
	"english_bot_admin/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	InsertForNewTask(ctx context.Context, taskId uuid.UUID, a string, b string, c string) error
	UpdateForTask(ctx context.Context, taskId uuid.UUID, answers *models.IncorrectAnswers) error
	SelectAnswersForTask(ctx context.Context, taskId uuid.UUID) (*models.IncorrectAnswers, error)
}
