package repo

type Task struct { // создали сущность Task
	ID          int    `json:"id"`          // id задачи
	Title       string `json:"title"`       // название задачи
	Description string `json:"description"` // описание
	Status      string `json:"status"`      // статус (например, "завершена", "в процессе" и тд)
}
