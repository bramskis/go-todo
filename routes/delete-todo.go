package routes

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTodo(c *gin.Context) {
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

	_, err = db.Exec("DELETE FROM todo WHERE id = $1;", id)
	if err != nil {
		if DEBUG_MODE == "true" {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"Error": fmt.Sprintf("Unable to delete todo: %s", err.Error()),
				},
			)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
