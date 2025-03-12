package service

import (
	"fiber-service/internal/models"
	"fiber-service/internal/repo"
	"go.uber.org/zap"
)

// TaskService предоставляет бизнес-логику для работы с задачами
type TaskService struct {
	repo   repo.Task
	logger *zap.SugaredLogger
}

// NewTaskService создает новый сервис для работы с задачами
func NewTaskService(repo repo.Task, logger *zap.SugaredLogger) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

// Создание задачи
func (s *TaskService) CreateTask(task models.Task) (int, error) {
	// Дополнительная бизнес-логика (например, проверка полей задачи)
	return s.repo.CreateTask(task)
}

// Получение всех задач
func (s *TaskService) GetAllTasks() (map[int]models.Task, error) {
	return s.repo.GetAllTasks()
}

// Получение задачи по ID
func (s *TaskService) GetTaskById(id int) (models.Task, error) {
	return s.repo.GetTaskById(id)
}

// Обновление задачи по ID
func (s *TaskService) UpdateTaskById(id int, task models.Task) error {
	// Дополнительная бизнес-логика (например, валидация статуса)
	return s.repo.UpdateTaskById(id, task)
}

// Удаление задачи по ID
func (s *TaskService) DeleteTaskById(id int) error {
	return s.repo.DeleteTaskById(id)
}

// Service — основная структура, включающая TaskService
type Service struct {
	Task   *TaskService
	logger *zap.SugaredLogger
}

// NewService — конструктор для создания экземпляра Service
func NewService(repo *repo.TaskRepository, logger *zap.SugaredLogger) *Service {
	return &Service{
		Task:   NewTaskService(repo, logger), // Создаем экземпляр TaskService
		logger: logger,
	}
}
