package http

import (
	"bytes"
	"english_bot_admin/internal/incorrect"
	"english_bot_admin/internal/incorrect/models"
	"english_bot_admin/internal/task"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"html/template"
	"log"
	"strconv"
)

type TaskHandler struct {
	useCase        task.Usecase
	taskRepo       task.Repository
	incAnswersRepo incorrect.Repository
}

func NewTaskHandler(taskUseCase task.Usecase, taskRepo task.Repository, incRepo incorrect.Repository) *TaskHandler {
	return &TaskHandler{
		useCase:        taskUseCase,
		taskRepo:       taskRepo,
		incAnswersRepo: incRepo,
	}
}

func (h *TaskHandler) RenderTasks(ctx *fiber.Ctx, tasks []task.Task) {
	tmpl, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Tasks []task.Task
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
	tasks, err := h.useCase.GetTasks(context_)
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
		return err
	}

	h.RenderTasks(ctx, tasks)
	return nil
}

func (h *TaskHandler) EditTask(ctx *fiber.Ctx) error {
	var errorMessage string = "Task updated successfully"
	context_ := ctx.Context()
	uuid_ := ctx.FormValue("task_id")
	taskUuid, err := uuid.Parse(uuid_)
	if err != nil {
		errorMessage = "Parsing uuid error"
		return err
	}
	taskType := ctx.FormValue("type")
	level := ctx.FormValue("level")
	question := ctx.FormValue("question")
	answer := ctx.FormValue("answer")
	taskTypeInt, err := strconv.Atoi(taskType)
	if err != nil {
		errorMessage = "Error task type converting to integer"
	}
	editTask := &task.Task{
		TaskID:   taskUuid,
		TypeID:   uint8(taskTypeInt),
		Level:    level,
		Question: question,
		Answer:   answer,
	}

	err = h.useCase.UpdateTaskInfoByUUID(context_, editTask)
	if err != nil {
		errorMessage = "Updating task info error"
	}

	incAnswerA := ctx.FormValue("incorrectA")
	incAnswerB := ctx.FormValue("incorrectB")
	incAnswerC := ctx.FormValue("incorrectC")

	incAnswers := &models.IncorrectAnswers{
		A: incAnswerA,
		B: incAnswerB,
		C: incAnswerC,
	}

	// TODO: 3

	err = h.incAnswersRepo.UpdateForTask(context_, taskUuid, incAnswers)
	if err != nil {
		errorMessage = "Updating of incorrect answers error"
	}

	data := fiber.Map{
		"Message": errorMessage,
	}
	return ctx.Render("templates/message.html", data)
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var internalId uuid.UUID
	context_ := ctx.Context()
	taskType := ctx.FormValue("type")
	level := ctx.FormValue("level")
	question := ctx.FormValue("question")
	answer := ctx.FormValue("answer")

	taskTypeInt, err := strconv.Atoi(taskType)

	newTask := &task.Task{
		TaskID:   uuid.New(),
		TypeID:   uint8(taskTypeInt),
		Level:    level,
		Question: question,
		Answer:   answer,
	}

	internalId, err = h.useCase.CreateTask(context_, newTask)
	if err != nil {
		return err
	}
	if err != nil {
		return nil
	}

	incAnswerA := ctx.FormValue("incorrectA")
	incAnswerB := ctx.FormValue("incorrectB")
	incAnswerC := ctx.FormValue("incorrectC")

	// TODO: 3

	err = h.incAnswersRepo.AddForNewTask(context_, internalId, incAnswerA, incAnswerB, incAnswerC)
	if err != nil {
		return err
	}

	return ctx.SendString("Task created successfully!")
}

func (h *TaskHandler) GetNewTaskForm(ctx *fiber.Ctx) error {
	return ctx.Render("templates/create_task.html", fiber.Map{})
}

func (h *TaskHandler) GetEditTaskForm(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	id_ := ctx.Params("id")
	uuid_, err := uuid.Parse(id_)
	if err != nil {
		return err
	}
	task_, err := h.useCase.GetTaskById(context_, uuid_)
	if err != nil {
		return err
	}
	if task_ == nil {
		log.Panic("задачи для редактирования нет")
	}

	// TODO: 4

	insAnswers, err := h.incAnswersRepo.GetAnswersForTask(context_, uuid_)
	if err != nil {
		log.Println("Error getting incorrect answers")
		return err
	}
	data := fiber.Map{
		"TaskID":     uuid_,
		"Task":       task_,
		"TypeID":     task_.TypeID,
		"Level":      task_.Level,
		"Question":   task_.Question,
		"Answer":     task_.Answer,
		"IncorrectA": insAnswers.A,
		"IncorrectB": insAnswers.B,
		"IncorrectC": insAnswers.C,
	}
	return ctx.Render("templates/edit_task.html", data)
}

func (h *TaskHandler) AddToModule(ctx *fiber.Ctx) error {
	//context_ := ctx.Context()
	//moduleID := ctx.Params("module_id")
	//taskID := ctx.Params("task_id")
	return nil
}

func (h *TaskHandler) BaseView(ctx *fiber.Ctx) error {
	return ctx.Render("templates/base.html", fiber.Map{})
}

func (h *TaskHandler) GetTaskAccess(ctx *fiber.Ctx) error {
	// отображение всех задач но добавить кнопки с галочкой
	return nil
}

func (h *TaskHandler) EditTaskAccess(ctx *fiber.Ctx) error {
	// нажать кнопку сохранить после проставления галочек из get
	// редактирование происходит для определенного пользоватлея? модуля? разбить на модуль
	return nil
}
