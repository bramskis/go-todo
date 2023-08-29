package main

import (
	"bramskis/go-todo/routes"
	"bramskis/go-todo/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.ContextWithFallback = true

	initializeDBTable()

	r.GET("/", routes.Index)
	r.GET("/todo", routes.GetTodoAll)
	r.GET("/todo/:id", routes.GetTodo)
	r.POST("/todo", routes.CreateTodo)
	r.PUT("/todo", routes.UpdateTodo)
	r.DELETE("/todo/:id", routes.DeleteTodo)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func initializeDBTable() {
	db, err := utils.GetDBConnection()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS todo (
			id VARCHAR(255) NOT NULL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			deadline TIMESTAMP NOT NULL,
			completed BOOLEAN NOT NULL
		);`,
	)
	if err != nil {
		panic(err)
	}
}
