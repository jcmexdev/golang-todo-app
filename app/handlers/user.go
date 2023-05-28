package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jxmexdev/go-todo-app/app/models"
	"github.com/jxmexdev/go-todo-app/app/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UpdateUserRequest struct {
	Email string `json:"email"`
}

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		existUser, err := repository.FindUserByEmail(r.Context(), user.Email)

		if existUser != nil && err == nil {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		if err != nil && err != mongo.ErrNoDocuments {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err = repository.CreateUser(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func FindAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.FindAllUsers(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			return
		}
	}
}

func FindUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user, err := repository.FindUserById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request UpdateUserRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		params := mux.Vars(r)
		user, err := repository.FindUserById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.Email = request.Email
		user, err = repository.UpdateUser(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
