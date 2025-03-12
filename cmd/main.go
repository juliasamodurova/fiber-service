package main

import (
	"fiber-service/internal/config"              // импорт конфигураций
	customLogger "fiber-service/internal/logger" // импортируем логгер
	"fiber-service/internal/repo"                // импортируем пакет, где будет храниться наша сущность Task
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"strconv"
)

// глобальные переменные для хранения данных
var tasks = make(map[int]repo.Task)
var lastID = 0 // переменная для генерации уникальных ID

func main() {

	// Загружаем конфигурацию из переменных окружения
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("Ошибка загрузки env файла:", err)
	}

	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	// Инициализация логгера
	logger, err := customLogger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal("Ошибка инициализации логгера:", err)
	}

	app := fiber.New()

	// корневой маршрут, проверяем, запущен ли сервер
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// создаём метод POST /tasks - "создание новой задачи"
	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task repo.Task
		if err := c.BodyParser(&task); err != nil {
			logger.Warnf("Ошибка при разборе тела запроса: %v", err) // логирование ошибки
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		lastID++              // увеличиваем ID на 1
		task.ID = lastID      // присваиваем новый ID
		tasks[task.ID] = task // сохраняем в map
		// логирование успешного создания задачи
		logger.Infof("Задача с ID %d успешно создана", task.ID)
		return c.JSON(task)
	})

	// создаём метод GET /tasks - "получение списка всех задач"
	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.JSON(tasks)
	})

	// создаём метод GET /tasks/:id – "получение задачи по ID"
	app.Get("/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id")) // получаем ID задачи из URL-параметра (/tasks/1 → "1") и пытаемся преобразовать его в int
		if err != nil {                         // если ID не является числом, возвращаем ошибку 400 (Bad Request)
			logger.Warnf("Некорректный ID задачи: %v", err) // логирование ошибки
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}

		task, exist := tasks[id] // ищем задачу в карте `tasks` по ID
		if !exist {              // если задача не найдена, возвращаем 404 (Not Found)
			logger.Warnf("Задача с ID %d не найдена", id) // логирование ошибки
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}

		return c.JSON(task) // если задача найдена, отправляем её в ответе
	})

	// создаём метод PUT /tasks/:id – "обновление задачи"
	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		// получаем ID задачи из URL-параметра и пытаемся преобразовать его в int
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			logger.Warnf("Некорректный ID задачи: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}
		// ищем задачу в map `tasks` по ID и сохраняем саму задачу в переменную task
		task, exist := tasks[id]
		if !exist {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		// парсим тело запроса и записываем данные в `task`:
		if err := c.BodyParser(&task); err != nil {
			logger.Warnf("Ошибка при разборе тела запроса: %v", err) // логирование парсинга
			// если произошла ошибка при парсинге JSON, возвращаем 400 (Bad Request)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		task.ID = id     // сохраняем прежний ID, чтобы он не изменился при обновлении
		tasks[id] = task // заменяем старую версию задачи новой в map

		return c.JSON(task)
	})
	// добавляем новый метод DELETE /tasks/:id – "удаление задачи"
	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			logger.Warnf("Некорректный ID задачи: %v", err) // логирование ошибки
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}
		// не планируем работать с содержимым задачи, нам важно только наличие её в коллекции.
		// поэтому в переменную task не сохраняем:
		_, exists := tasks[id]
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		delete(tasks, id)                                  // удаляем задачу
		logger.Infof("Задача с ID %d успешно удалена", id) // логирование успешного удаления
		return c.SendString("Task deleted")
	})

	// запуск сервера на порту 8080
	logger.Debug("Сервер запущен на порту :8080")
	if err := app.Listen(":8080"); err != nil {
		logger.Fatal("Ошибка запуска сервера:", err)
	}
}
