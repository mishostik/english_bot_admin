package httpserver

import (
	"english_bot_admin/database"
	incAnswersRepository "english_bot_admin/internal/incorrect/repository"
	"english_bot_admin/internal/module/http"
	moduleRepository "english_bot_admin/internal/module/repository"
	moduleUseCase "english_bot_admin/internal/module/usecase"
	taskHttp "english_bot_admin/internal/task/http"
	taskRepository "english_bot_admin/internal/task/repository"
	taskUseCase "english_bot_admin/internal/task/usecase"
	"log"
)

func MapHandlers(db *database.Database, s *Server) error {

	taskCollection, err := db.Collection(TasksCollection)
	if err != nil {
		log.Fatalf("error connection {tasks}: %v", err.Error())
	}
	typeCollection, err := db.Collection(TypesCollection)
	if err != nil {
		log.Fatalf("error connection {task types}: %v", err.Error())
	}
	incorrectAnswers, err := db.Collection(IncAnswersCollection)
	if err != nil {
		log.Fatalf("error connection {incorrect answers}: %v", err.Error())
	}
	moduleCollection, err := db.Collection(ModulesCollection)
	if err != nil {
		log.Fatalf("error connection {modules}: %v", err.Error())
	}

	// ------------------------ repositories ------------------------

	moduleRepo := moduleRepository.NewModuleRepository(moduleCollection)
	taskRepo := taskRepository.NewMongoTaskRepository(taskCollection, typeCollection)
	incAnswersRepo := incAnswersRepository.NewIncorrectRepository(incorrectAnswers)

	// ------------------------- use cases -------------------------

	moduleUC := moduleUseCase.NewModuleUsecase(*moduleRepo)
	taskUC := taskUseCase.NewTaskUsecase(moduleUC, taskRepo)

	taskHandler := taskHttp.NewTaskHandler(taskUC, taskRepo, incAnswersRepo)

	taskHttp.TaskRoutes(s.app, taskHandler)

	moduleHandler := http.NewModuleHandler(moduleUC)
	http.ModuleRoutes(s.app, moduleHandler)

	return nil

}
