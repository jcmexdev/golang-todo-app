package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jxmexdev/go-todo-app/app/models"
	"github.com/jxmexdev/go-todo-app/app/repository"
	"net/http"
)

func CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task *models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task, err = repository.CreateTask(r.Context(), task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

type UpdateTaskRequest struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func UpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var taskRequest UpdateTaskRequest
		err := json.NewDecoder(r.Body).Decode(&taskRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, err := repository.FindTaskById(r.Context(), vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		task.Description = taskRequest.Description
		task.Completed = taskRequest.Completed

		task, err = repository.UpdateTask(r.Context(), task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func FindTaskById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		task, err := repository.FindTaskById(r.Context(), vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func FindAllTasksByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tasks, err := repository.FindAllTasksByUserId(r.Context(), vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func DeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		err := repository.DeleteTask(r.Context(), vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
