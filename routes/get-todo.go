package routes

import (
	"bramskis/go-todo/types"
	"bramskis/go-todo/utils"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func GetTodo(c *gin.Context) {
	id := c.Param("id")

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

	var toReturn types.Todo
	if err = db.QueryRow("SELECT * FROM todo WHERE id = $1;", id).Scan(
		&toReturn.Id, &toReturn.Title, &toReturn.Description, &toReturn.Deadline, &toReturn.Completed,
	); err != nil {
		fmt.Println("hit an error")
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
		return
	}

	c.JSON(http.StatusOK, toReturn)
}
