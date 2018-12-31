package todolist_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app"
	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app/todolist"
)

type MockService struct{}

func (this MockService) FindAll() ([]todolist.Task, error) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2018-12-31T15:00:00.300Z"
	t, _ := time.Parse(layout, str)
	todos := []todolist.Task{
		{ID: 1, Description: "Item 1", CreatedAt: t, Done: false},
		{ID: 2, Description: "Item 2", CreatedAt: t, Done: true},
	}
	return todos, nil
}

func (this MockService) Create(todo todolist.Task) (todolist.Task, error) {
	todos, err := this.FindAll()
	return todos[0], err
}

func (this MockService) Done(todo todolist.Task) (todolist.Task, error) {
	todos, err := this.FindAll()
	return todos[1], err
}

var application *app.Application

func setUpHandler() {
	handler := todolist.NewHandler(MockService{})
	application = app.New()
	application.Handler = &handler
	application.Initialize(map[string]string{
		"databaseURL": "",
	})
}

func TestFindAll(t *testing.T) {
	setUpHandler()
	req, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	application.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	resp := rr.Body.String()
	results := []todolist.Task{}
	json.Unmarshal([]byte(resp), &results)
	expected := 2
	if len(results) != expected {
		t.Errorf("handler returned unexpected body: got %d want %d", len(results), expected)
	}
}

func TestCreate(t *testing.T) {
	setUpHandler()
	var jsonStr = []byte(`{"description": "Item 2"}`)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	application.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	resp := rr.Body.String()
	result := todolist.Task{}
	json.Unmarshal([]byte(resp), &result)
	expected := 1
	if result.ID != expected {
		t.Errorf("handler returned a wrong ID: got %d want %d", result.ID, expected)
	}
}

func TestDone(t *testing.T) {
	setUpHandler()
	req, _ := http.NewRequest("PUT", "/todos/2", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	application.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	resp := rr.Body.String()
	result := todolist.Task{}
	json.Unmarshal([]byte(resp), &result)
	expected := 2
	if result.ID != expected {
		t.Errorf("handler returned a wrong ID: got %d want %d", result.ID, expected)
	}
	if !result.Done {
		t.Errorf("handler returned not completed: got done = %v want true", result.Done)
	}
}
