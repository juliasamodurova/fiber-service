package api

import (
	"fiber-service/internal/models"
	"fiber-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// API структура, которая будет хранить сервис и логгер
type API struct {
	TaskService *service.TaskService
	Logger      *zap.SugaredLogger
}

// NewAPI конструктор для создания нового API
func NewAPI(taskService *service.TaskService, logger *zap.SugaredLogger) *API {
	return &API{
		TaskService: taskService,
		Logger:      logger,
	}
}

// SetRoutes настраивает маршруты API
func (api *API) SetRoutes(app *fiber.App) {
	app.Post("/tasks", api.CreateTask)
	app.Get("/tasks", api.GetAllTasks)
	app.Get("/tasks/:id", api.GetTaskById)
	app.Put("/tasks/:id", api.UpdateTaskById)
	app.Delete("/tasks/:id", api.DeleteTaskById)
}

// далее очень много букаф для настройки обработчиков и логирования

// CreateTask обработчик для создания задачи
func (api *API) CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		api.Logger.Error("Error parsing task body", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// вызов метода сервиса для создания задачи
	taskID, err := api.TaskService.CreateTask(task)
	if err != nil {
		api.Logger.Error("Error creating task", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"id":      taskID,
	})
}

// GetAllTasks обработчик для получения всех задач
func (api *API) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := api.TaskService.GetAllTasks()
	if err != nil {
		api.Logger.Error("Error fetching all tasks", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tasks",
		})
	}

	return c.Status(http.StatusOK).JSON(tasks)
}

// GetTaskById обработчик для получения задачи по ID
func (api *API) GetTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		api.Logger.Error("Invalid task ID", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}

	task, err := api.TaskService.GetTaskById(id)
	if err != nil {
		api.Logger.Error("Error fetching task by ID", zap.Error(err))
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	return c.Status(http.StatusOK).JSON(task)
}

// UpdateTaskById обработчик для обновления задачи по ID
func (api *API) UpdateTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		api.Logger.Error("Invalid task ID", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		api.Logger.Error("Error parsing task body", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	err = api.TaskService.UpdateTaskById(id, task)
	if err != nil {
		api.Logger.Error("Error updating task", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
	})
}

// DeleteTaskById обработчик для удаления задачи по ID
func (api *API) DeleteTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		api.Logger.Error("Invalid task ID", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}

	err = api.TaskService.DeleteTaskById(id)
	if err != nil {
		api.Logger.Error("Error deleting task", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}

// Routers структура для работы с маршрутизацией
type Routers struct {
	Service *service.TaskService
	Logger  *zap.SugaredLogger
}

// NewRouters конструктор для создания экземпляра Routers
func NewRouters(service *service.TaskService, logger *zap.SugaredLogger) *Routers {
	return &Routers{
		Service: service,
		Logger:  logger,
	}
}

// SetRoutes настраивает маршруты для приложения
func (r *Routers) SetRoutes(app *fiber.App) {
	api := NewAPI(r.Service, r.Logger) // создаем API с нужными сервисами
	api.SetRoutes(app)                 // настраиваем маршруты для API
}

// Listen метод для запуска сервера
func (r *Routers) Listen(addr string) error {
	// Создаем новое приложение Fiber
	app := fiber.New()

	// Настроим маршруты с помощью SetRoutes
	r.SetRoutes(app)

	// Запускаем сервер
	return app.Listen(addr)
}
