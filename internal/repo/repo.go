package repo

import (
	"fiber-service/internal/models"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

type Task interface {
	CreateTask(task models.Task) (int, error)
	GetAllTasks() (map[int]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	UpdateTaskById(id int, task models.Task) error
	DeleteTaskById(id int) error
}

type TaskRepository struct {
	mu     *sync.RWMutex
	tasks  map[int]models.Task
	nextID int
	logger *zap.SugaredLogger
}

// конструктор для создания нового TaskRepository
func NewTaskRepository(mu *sync.RWMutex, logger *zap.SugaredLogger) Task {
	taskRepo := &TaskRepository{
		mu:     mu,
		tasks:  make(map[int]models.Task),
		logger: logger,
	}

	// указываем, что taskRepo реализует интерфейс Task
	return taskRepo
}

// репозиторий, который использует Task
type Repository struct {
	Task Task
}

// конструктор для создания нового репозитория с Task
func NewRepository(mu *sync.RWMutex, logger *zap.SugaredLogger) *TaskRepository {
	return &TaskRepository{
		mu:     mu,
		tasks:  make(map[int]models.Task),
		logger: logger,
	}
}
func (r *TaskRepository) CreateTask(task models.Task) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.nextID++ // увеличиваем nextID
	task.ID = r.nextID
	r.tasks[r.nextID] = task

	r.logger.Info("Task created", zap.Int("taskID", task.ID)) // логируем создание задачи
	return task.ID, nil
}

func (r *TaskRepository) GetAllTasks() (map[int]models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Info("Fetching all tasks") // логируем получение всех задач
	return r.tasks, nil
}

func (r *TaskRepository) GetTaskById(id int) (models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		r.logger.Warn("Task not found", zap.Int("taskID", id)) // логируем предупреждение, если задача не найдена
		return models.Task{}, fmt.Errorf("task with id %d not found", id)
	}

	r.logger.Info("Task fetched", zap.Int("taskID", id)) // логируем получение задачи по ID
	return task, nil
}

func (r *TaskRepository) UpdateTaskById(id int, task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.tasks[id]
	if !exists {
		r.logger.Warn("Task not found for update", zap.Int("taskID", id)) // логируем предупреждение, если задача не найдена
		return fmt.Errorf("task with id %d not found", id)
	}

	r.tasks[id] = task
	r.logger.Info("Task updated", zap.Int("taskID", id)) // логируем успешное обновление задачи
	return nil
}

func (r *TaskRepository) DeleteTaskById(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.tasks[id]
	if !exists {
		r.logger.Warn("Task not found for deletion", zap.Int("taskID", id)) // логируем предупреждение, если задача не найдена
		return fmt.Errorf("task with id %d not found", id)
	}

	delete(r.tasks, id)
	r.logger.Info("Task deleted", zap.Int("taskID", id)) // логируем успешное удаление задачи
	return nil
}
