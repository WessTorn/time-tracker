package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

func stopTask(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(stopTask)")
	getId := c.Param("id")

	var request struct {
		TaskID int `json:"task_id"`
	}

	id, err := strconv.Atoi(getId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	task := database.Task{
		UserID:  id,
		TaskID:  request.TaskID,
		EndTime: time.Now(),
	}

	if !database.IsTaskStarted(db, task) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not started"})
		return
	}

	err = database.UpdateTask(db, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task stopped successfully"})
}
