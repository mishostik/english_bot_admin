package incorrect

import (
	"context"
	"english_bot_admin/internal/incorrect/models"
	"github.com/google/uuid"
)

type Repository interface {
	AddForNewTask(ctx context.Context, taskId uuid.UUID, a string, b string, c string) error
	UpdateForTask(ctx context.Context, taskId uuid.UUID, answers *models.IncorrectAnswers) error
	GetAnswersForTask(ctx context.Context, taskId uuid.UUID) (*models.IncorrectAnswers, error)
}
