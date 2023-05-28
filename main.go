package main

import (
	"fmt"
	"github.com/jxmexdev/go-todo-app/app/env"
	"github.com/jxmexdev/go-todo-app/app/repository"
	"github.com/jxmexdev/go-todo-app/app/router"
	"log"
	"net/http"
)

func main() {
	env.LoadConfiguration()
	repository.LoadConfiguration()
	r := router.GetRouter()
	uri := env.Conf.AppHost + ":" + env.Conf.AppPort
	fmt.Println("Server running on port", uri)
	log.Fatal(http.ListenAndServe(uri, r))
}
