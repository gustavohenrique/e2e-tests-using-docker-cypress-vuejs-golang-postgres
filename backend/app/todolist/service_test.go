package todolist_test

import (
	"log"
	"testing"

	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app/database"
	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app/todolist"
)

var (
	db      database.Database
	service todolist.TaskService
)

func setUpService() {
	db = database.NewDB("postgres://admin:password@127.0.0.1:5434/todos?sslmode=disable")
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database test.", err)
	}
	conn.Exec("TRUNCATE tasks RESTART IDENTITY;")
	tx, _ := conn.Begin()
	stmt, _ := tx.Prepare("INSERT INTO tasks (description, done) VALUES ($1, $2)")
	stmt.Exec("Write a new post", false)
	stmt.Exec("Fix some bugs in my app", false)
	stmt.Exec("Install the new Linux Mint version", false)
	stmt.Exec("Create a TASK app as example of e2e testing", true)
	tx.Commit()

	service = todolist.NewService(db)
}

func TestFindAllShouldReturnsAllTasks(t *testing.T) {
	setUpService()
	tasks, _ := service.FindAll()
	expected := 4
	if len(tasks) != expected {
		t.Errorf("Oops, failed to find all tasks. Got %d but I want %d", len(tasks), expected)
	}
}

func TestCreateNewTaskShouldReturnsItWithID(t *testing.T) {
	setUpService()
	task := todolist.Task{
		Description: "Testing create todo",
	}
	saved, _ := service.Create(task)
	expected := 5
	if expected != saved.ID {
		t.Errorf("The new todo ID is %d but the expected is %d", saved.ID, expected)
	}
}

func TestUpdateTaskAsDoneShouldReturnsItWithDoneEqualTrue(t *testing.T) {
	setUpService()
	task := todolist.Task{
		ID: 3,
	}
	saved, _ := service.Done(task)
	if !saved.Done {
		t.Fail()
	}
}
