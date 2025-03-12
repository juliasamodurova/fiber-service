## **Описание проекта**

Fiber Service – это REST API-сервис, написанный на Go с использованием фреймворка Fiber. Он позволяет создавать, читать, обновлять и удалять задачи. Задачи хранятся в памяти, но проект
служит основой для создания более сложного сервиса в будущем.

## Технологии
- Go
- Fiber
- envconfig (для управления конфигурациями)
- логирование с использованием `zap`

---
## Как запустить
1. Клонируйте этот репозиторий:
   ```bash
   git clone https://github.com/вашusername/fiber-service.git
   ```
2. Установите зависимости:
   ```bash
   go mod tidy
   ```
3. Настройте файл .env: Убедитесь, что файл .env присутствует в проекте и содержит правильные настройки:
   ```bash
   PORT=:8080
   SERVER_NAME=FiberService
   ```
4. Запустите сервер:
   ```bash
   go run cmd/main.go
   ```

Сервер будет запущен на порту 8080 по умолчанию.

Далее вы можете взаимодействовать с API через Postman или любой другой HTTP клиент.

---
## API Эндпоинты
- `POST /tasks` - Создать новую задачу
Переключитесь на метод POST.
Введите URL: http://localhost:8080/tasks.
Перейди на вкладку Body и выбери raw. Вставь следующий JSON в тело запроса:
```
{
  "title": "New Task",
  "description": "This is a new task",
  "status": "in progress"
}
```
Нажми Send. 
Ответ :
```
{
"id": 1,
"title": "New Task",
"description": "This is a new task",
"status": "in progress"
}
```
- `GET /tasks` - Получить все задачи
  http://localhost:8080/tasks
Ответ:
```
{
  "1": {
    "id": 1,
    "title": "New Task",
    "description": "This is a new task",
    "status": "in progress"
  }
}
```
- `GET /tasks/:id` - Получить задачу по ID
  http://localhost:8080/tasks/1
Запрос:
```
{
  "id": 1,
  "title": "New Task",
  "description": "This is a new task",
  "status": "in progress"
}
```
Ответ:
```
{
"id": 1,
"title": "New Task",
"description": "This is a new task",
"status": "in progress"
}
```

- `PUT /tasks/:id` - Обновить задачу по ID
  http://localhost:8080/tasks/1
Запрос:
```
{
  "title": "Updated Task",
  "description": "This task has been updated",
  "status": "completed"
}
```
Ответ:
```
{
"id": 1,
"title": "Updated Task",
"description": "This task has been updated",
"status": "completed"
}
```
- `DELETE /tasks/:id` - Удалить задачу по ID
  http://localhost:8080/tasks/1
  Send
  Ответ: "Task deleted"

---

Сервис готов к работе.