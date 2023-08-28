package routes

import (
	"bramskis/go-todo/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTodo(c *gin.Context) {
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

	_, err = db.Exec("DELETE FROM todo WHERE id = $1;", id)
	if err != nil {
		if utils.IsDebugMode() {
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
