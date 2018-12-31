package todolist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) TaskHandler {
	handler := TaskHandler{}
	handler.service = service
	return handler
}

func (this TaskHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := this.service.FindAll()
	if err != nil {
		log.Println("#FindAll:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (this TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var task Task
	err = json.Unmarshal(b, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Creating task %v", task)
	saved, err := this.service.Create(task)
	if err != nil {
		log.Println("#Create:", err)
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(saved)
}

func (this TaskHandler) Done(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	task := Task{ID: id}
	saved, err := this.service.Done(task)
	if err != nil {
		log.Println("#Done:", err)
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(saved)
}
