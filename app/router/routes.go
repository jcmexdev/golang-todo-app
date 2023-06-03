package router

import (
	"github.com/gorilla/mux"
	"github.com/jxmexdev/go-todo-app/app/handlers"
	"net/http"
)

func GetRouter() (router *mux.Router) {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.FindAllUsers()).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.CreateUser()).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", handlers.FindUserById()).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handlers.UpdateUser()).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}/tasks", handlers.FindAllTasksByUserId()).Methods(http.MethodGet)
	r.HandleFunc("/tasks", handlers.CreateTask()).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask()).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", handlers.FindTaskById()).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask()).Methods(http.MethodDelete)
	return r
}
