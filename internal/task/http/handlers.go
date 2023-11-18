package http

import (
	"bytes"
	incr "english_bot_admin/internal/incorrect/repository"
	model "english_bot_admin/internal/task/models"
	tr "english_bot_admin/internal/task/repository"
	uc "english_bot_admin/internal/task/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"strconv"
)

type TaskHandler struct {
	useCase        *uc.TaskUsecase
	taskRepo       *tr.MongoTaskRepository
	incAnswersRepo *incr.IncorrectRepository
}

func NewTaskHandler(taskCollection *mongo.Collection, typeCollection *mongo.Collection, incAnswers *mongo.Collection) *TaskHandler {
	return &TaskHandler{
		useCase:        uc.NewTaskUsecase(),
		taskRepo:       tr.NewMongoTaskRepository(taskCollection, typeCollection),
		incAnswersRepo: incr.NewIncorrectRepository(incAnswers),
	}
}

func RenderTasks(ctx *fiber.Ctx, tasks []model.Task) {
	tmpl, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Tasks []model.Task
	}{
		Tasks: tasks,
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}
	ctx.Set("Content-Type", "text/html")
	err = ctx.Status(fiber.StatusOK).Send(buf.Bytes())
	if err != nil {
		return
	}
}

func (h *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	tasks, err := h.taskRepo.GetTasks(context_)
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
		return err
	}

	RenderTasks(ctx, tasks)
	return nil
}

func (h *TaskHandler) EditTask(ctx *fiber.Ctx) error {
	log.Println("editing...")
	context_ := ctx.Context()
	id_ := ctx.Params("id")
	uuid_, err := uuid.Parse(id_)
	log.Println("task uuid -", uuid_)
	if err != nil {
		return err
	}
	task, err := h.taskRepo.GetTaskByID(context_, uuid_)
	if err != nil {
		log.Println("задачи нет по uuid")
		return err
	}
	if task == nil {
		log.Panic("задачи для редактирования нет")
	}
	insAnswers, err := h.incAnswersRepo.GetAnswersForTask(context_, uuid_)
	if err != nil {
		log.Println("Error while getting incorrect answers")
		return err
	}
	data := fiber.Map{
		"TaskID":     uuid_,
		"Task":       task,
		"TypeID":     task.TypeID,
		"Level":      task.Level,
		"Question":   task.Question,
		"Answer":     task.Answer,
		"IncorrectA": insAnswers.A,
		"IncorrectB": insAnswers.B,
		"IncorrectC": insAnswers.C,
	}
	log.Println("data from db about editing task -", data)
	return ctx.Render("templates/edit_task.html", data)
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var internalId uuid.UUID
	context_ := ctx.Context()
	taskType := ctx.FormValue("type")
	level := ctx.FormValue("level")
	question := ctx.FormValue("question")
	answer := ctx.FormValue("answer")

	taskTypeInt, err := strconv.Atoi(taskType)

	newTask := &model.Task{
		TaskID:   uuid.New(),
		TypeID:   uint8(taskTypeInt),
		Level:    level,
		Question: question,
		Answer:   answer,
	}
	internalId, err = h.taskRepo.NewTask(context_, newTask)
	if err != nil {
		return err
	}
	if err != nil {
		return nil
	}

	incAnswerA := ctx.FormValue("incorrectA")
	incAnswerB := ctx.FormValue("incorrectB")
	incAnswerC := ctx.FormValue("incorrectC")

	err = h.incAnswersRepo.AddForNewTask(context_, internalId, incAnswerA, incAnswerB, incAnswerC)
	if err != nil {
		return err
	}

	return ctx.SendString("Task created successfully!")
}
