package routes

import (
	"bramskis/go-todo/types"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var input types.CreateTodoRequest
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

	sqlStatement := "INSERT INTO todo(id, title, description, deadline, completed) VALUES( $1, $2, $3, $4, $5 )"

	_, err = db.Exec(sqlStatement, uuid.New().String(), input.Title, input.Description, input.Deadline, input.Completed)
	if err != nil {
		if DEBUG_MODE == "true" {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"Error": fmt.Sprintf("Unable to create todo: %s", err.Error()),
				},
			)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
