package routes

import (
	"bramskis/go-todo/types"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func GetTodo(c *gin.Context) {
	id := c.Param("id")

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

	var toReturn types.Todo
	if err = db.QueryRow("SELECT * FROM todo WHERE id = ?", id).Scan(
		&toReturn.Id, &toReturn.Title, &toReturn.Description, &toReturn.Deadline, &toReturn.Completed,
	); err != nil {
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
		return
	}

	c.JSON(http.StatusOK, toReturn)
}
