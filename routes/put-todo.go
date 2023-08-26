package routes

import (
	"bramskis/go-todo/types"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateTodo(c *gin.Context) {
	var input types.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Request must contain a valid Todo"})
		return
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		if DEBUG_MODE == "true" {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"Error": fmt.Sprintf("Db connection issue: %s", err.Error()),
				},
			)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		_ = c.Error(err)
	}

	sqlStatement := `
		UPDATE todo
		SET title = $2, description = $3, deadline = $4, completed = $5
		WHERE id = $1;`
	_, err = db.Exec(sqlStatement, input.Id, input.Title, input.Description, input.Deadline, input.Completed)
	if err != nil {
		if DEBUG_MODE == "true" {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"Error": fmt.Sprintf("Unable to update todo: %s", err.Error()),
				},
			)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
