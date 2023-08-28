package routes

import (
	"bramskis/go-todo/types"
	"bramskis/go-todo/utils"
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

	db, err := utils.GetDBConnection()
	if err != nil {
		if utils.IsDebugMode() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	sqlStatement := "INSERT INTO todo(id, title, description, deadline, completed) VALUES($1, $2, $3, $4, $5)"

	_, err = db.Exec(sqlStatement, uuid.New().String(), input.Title, input.Description, input.Deadline, input.Completed)
	if err != nil {
		if utils.IsDebugMode() {
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
