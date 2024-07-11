package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

func getTasks(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(getTasks)")

	getId := c.Param("id")

	id, err := strconv.Atoi(getId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := database.SelectTasks(db, id)
	if err != nil {
		switch err.Error() {
		case "TasksNotFound":
			c.JSON(http.StatusNotFound, gin.H{"error": "Tasks not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks"})
		}
		return
	}

	c.JSON(http.StatusOK, tasks)
}
