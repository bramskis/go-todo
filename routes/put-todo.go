package routes

import (
	"bramskis/go-todo/types"
	"bramskis/go-todo/utils"
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

	sqlStatement := `
		UPDATE todo
		SET title = $2, description = $3, deadline = $4, completed = $5
		WHERE id = $1;
	`
	_, err = db.Exec(sqlStatement, input.Id, input.Title, input.Description, input.Deadline, input.Completed)
	if err != nil {
		if utils.IsDebugMode() {
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
