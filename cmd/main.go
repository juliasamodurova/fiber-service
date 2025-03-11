package main

import (
	"fiber-service/internal/repo"
	"github.com/gofiber/fiber/v2"
	"log" // импортируем пакет, где будет храниться наша сущность Task
	"strconv"
)

// глобальные переменные для хранения данных
var tasks = make(map[int]repo.Task)
var lastID = 0 // переменная для генерации уникальных ID

func main() {
	app := fiber.New()

	// корневой маршрут, проверяем, запущен ли сервер
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// создаём метод POST /tasks - "создание новой задачи"
	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task repo.Task
		if err := c.BodyParser(&task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		lastID++              // увеличиваем ID на 1
		task.ID = lastID      // присваиваем новый ID
		tasks[task.ID] = task // сохраняем в map

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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}

		task, exist := tasks[id] // ищем задачу в карте `tasks` по ID
		if !exist {              // если задача не найдена, возвращаем 404 (Not Found)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}

		return c.JSON(task) // если задача найдена, отправляем её в ответе
	})

	// создаём метод PUT /tasks/:id – "обновление задачи"
	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		// получаем ID задачи из URL-параметра и пытаемся преобразовать его в int
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}
		// ищем задачу в map `tasks` по ID и сохраняем саму задачу в переменную task
		task, exist := tasks[id]
		if !exist {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		// парсим тело запроса и записываем данные в `task`:
		if err := c.BodyParser(&task); err != nil {
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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}
		// не планируем работать с содержимым задачи, нам важно только наличие её в коллекции.
		// поэтому в переменную task не сохраняем:
		_, exists := tasks[id]
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}

		return c.SendString("Task deleted")
	})

	// запуск сервера на порту 8080
	log.Fatal(app.Listen(":8080"))
}
