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
	return r
}
