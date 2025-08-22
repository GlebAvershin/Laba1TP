package main

import (
	"fmt"
	"todoapi/http"
	"todoapi/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandler(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server", err)
	}
}
