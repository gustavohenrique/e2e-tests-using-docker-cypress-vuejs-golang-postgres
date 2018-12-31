package todolist

import (
	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app/database"
)

type ServiceInterface interface {
	FindAll() ([]Task, error)
	Create(task Task) (Task, error)
	Done(task Task) (Task, error)
}

type TaskService struct {
	DB database.Database
}

func NewService(db database.Database) TaskService {
	service := TaskService{}
	service.DB = db
	return service
}

func (this TaskService) FindAll() ([]Task, error) {
	tasks := []Task{}
	db, err := this.DB.Connect()
	if err == nil {
		db.Select(&tasks, "SELECT * FROM tasks")
	}
	return tasks, err
}

func (this TaskService) Create(task Task) (Task, error) {
	db, err := this.DB.Connect()
	if err == nil {
		query := "INSERT INTO tasks (description) VALUES ($1) RETURNING *"
		stmt, err := db.Preparex(query)
		var saved Task
		err = stmt.Get(&saved, task.Description)
		return saved, err
	}
	return task, err
}

func (this TaskService) Done(task Task) (Task, error) {
	db, err := this.DB.Connect()
	if err == nil {
		query := "UPDATE tasks SET done = true WHERE id = $1 RETURNING *"
		stmt, err := db.Preparex(query)
		var saved Task
		err = stmt.Get(&saved, task.ID)
		return saved, err
	}
	return task, err
}
