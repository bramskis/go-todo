package main

import (
	"bramskis/go-todo/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.ContextWithFallback = true

	r.GET("/todo", routes.GetTodoAll)
	r.GET("/todo/:id", routes.GetTodo)
	r.POST("/todo", routes.CreateTodo)
	r.PUT("/todo", routes.UpdateTodo)
	r.DELETE("/todo/:id", routes.DeleteTodo)

	err := r.Run("0.0.0.0:3000")
	if err != nil {
		fmt.Printf("Run error encountered: %s", err.Error())
	}
}
