package usecase

import (
	"context"
	"english_bot_admin/internal/module"
	"english_bot_admin/internal/module/repository"
	"github.com/google/uuid"
	"log"
)

type ModuleUseCase struct {
	repo repository.ModuleRepository
}

func NewModuleUsecase(moduleRepo repository.ModuleRepository) module.Usecase {
	return &ModuleUseCase{repo: moduleRepo}
}

func (u *ModuleUseCase) GenerateModule(ctx context.Context, newModule module.NewModuleParams) error {
	temp := &module.Module{
		ModuleID: uuid.New(),
		Title:    newModule.Title,
		Level:    newModule.Level,
		Task:     newModule.Task,
	}
	err := u.repo.NewModule(ctx, temp)
	if err != nil {
		return err
	}
	return nil
}

func (u *ModuleUseCase) GetModules(ctx context.Context) ([]module.Module, error) {
	modules, err := u.repo.SelectModules(ctx)
	if err != nil {
		return []module.Module{}, err
	}
	if len(modules) == 0 {
		log.Println("modules amount is null")
	}
	return modules, nil
}

// ChangeModule todo: think about this method
func (u *ModuleUseCase) ChangeModule(newTasksNum []uuid.UUID) error {
	return nil
}

func (u *ModuleUseCase) GetModuleByID(moduleID uuid.UUID) (*module.Module, error) {
	return nil, nil
}

func (u *ModuleUseCase) AddTask(ctx context.Context, params module.TaskToModule) error {
	err := u.repo.InsertTask(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
