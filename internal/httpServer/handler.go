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
	userHttp "english_bot_admin/internal/user/http"
	userRepository "english_bot_admin/internal/user/repository"
	userUseCase "english_bot_admin/internal/user/usecase"

	"english_bot_admin/internal/httpServer/cconstants"
	"log"
)

func MapHandlers(db *database.Database, s *Server) error {

	taskCollection, err := db.Collection(cconstants.TasksCollection)
	if err != nil {
		log.Fatalf("error connection {tasks}: %v", err.Error())
	}
	typeCollection, err := db.Collection(cconstants.TypesCollection)
	if err != nil {
		log.Fatalf("error connection {task types}: %v", err.Error())
	}
	incorrectAnswers, err := db.Collection(cconstants.IncAnswersCollection)
	if err != nil {
		log.Fatalf("error connection {incorrect answers}: %v", err.Error())
	}
	moduleCollection, err := db.Collection(cconstants.ModulesCollection)
	if err != nil {
		log.Fatalf("error connection {modules}: %v", err.Error())
	}

	userCollection, err := db.Collection(cconstants.UsersCollection)

	// ------------------------ repositories ------------------------

	moduleRepo := moduleRepository.NewModuleRepository(moduleCollection)
	taskRepo := taskRepository.NewMongoTaskRepository(taskCollection, typeCollection)
	incAnswersRepo := incAnswersRepository.NewIncorrectRepository(incorrectAnswers)
	userRepo := userRepository.NewUserRepository(userCollection)

	// ------------------------- use cases -------------------------

	moduleUC := moduleUseCase.NewModuleUsecase(*moduleRepo)
	taskUC := taskUseCase.NewTaskUsecase(taskRepo)

	taskHandler := taskHttp.NewTaskHandler(taskUC, taskRepo, incAnswersRepo)
	taskHttp.TaskRoutes(s.app, taskHandler)

	moduleHandler := http.NewModuleHandler(moduleUC, taskUC)
	http.ModuleRoutes(s.app, moduleHandler)

	userUsecase := userUseCase.NewUserUsecase(userRepo)
	userHandler := userHttp.NewUserHandler(userUsecase)
	userHttp.UserRoutes(s.app, userHandler)

	return nil

}
