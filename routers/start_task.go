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

func startTask(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(startTask)")
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

	newTask := database.Task{
		UserID:    id,
		TaskID:    request.TaskID,
		StartTime: time.Now(),
	}

	if database.IsTaskStarted(db, newTask) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task already started"})
		return
	}

	err = database.InsertTask(db, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task started successfully"})
}
