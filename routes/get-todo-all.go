package routes

import (
	"bramskis/go-todo/types"
	"bramskis/go-todo/utils"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodoAll(c *gin.Context) {
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

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Request must contain a valid id"})
			return
		}
		if utils.IsDebugMode() {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"Error": fmt.Sprintf("Unable to retrieve todo: %s", err.Error()),
				},
			)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	var toReturn []types.Todo
	for rows.Next() {
		var todo types.Todo
		if err = rows.Scan(
			&todo.Id, &todo.Title, &todo.Description, &todo.Deadline, &todo.Completed,
		); err != nil {
			continue
		}
		toReturn = append(toReturn, todo)
	}

	c.JSON(http.StatusOK, toReturn)
}
