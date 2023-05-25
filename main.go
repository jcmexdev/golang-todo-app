package main

import (
	"github.com/jxmexdev/go-todo-app/app/env"
	"github.com/jxmexdev/go-todo-app/app/repository"
)

func main() {
	env.LoadConfiguration()
	repository.LoadConfiguration()
}
