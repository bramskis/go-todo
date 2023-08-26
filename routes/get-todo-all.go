package routes

import (
	"bramskis/go-todo/types"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodoAll(c *gin.Context) {
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

	rows, err := db.Query("SELECT * FROM todo")
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Request must contain a valid id"})
		return
	}
	if DEBUG_MODE == "true" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"Error": fmt.Sprintf("Unable to retrieve todo: %s", err.Error()),
			},
		)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var toReturn []types.Todo
	for rows.Next() {
		var todo types.Todo
		if err := rows.Scan(
			&todo.Id, &todo.Title, &todo.Description, &todo.Deadline, &todo.Completed,
		); err != nil {
			continue
		}
		toReturn = append(toReturn, todo)
	}

	c.JSON(http.StatusOK, toReturn)
}
