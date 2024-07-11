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

// @Summary Stop Task
// @Description Stops a task for a user
// @Tags tasks
// @Produce json
// @Param id path int true "User ID"
// @Param task_id body TaskID true "Task ID"
// @Success 200 {object} Response "Task stopped successfully"
// @Failure 400 {object} Response "Invalid user ID, Invalid request payload"
// @Failure 409 {object} Response "Task not started"
// @Failure 500 {object} Response "Failed to stop task"
// @Router /tasks/{id}/stop [post]
func stopTask(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(stopTask)")
	getId := c.Param("id")

	var request TaskID

	id, err := strconv.Atoi(getId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid user ID"})
		return
	}

	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid request payload"})
		return
	}

	task := database.Task{
		UserID:  id,
		TaskID:  request.TaskID,
		EndTime: time.Now(),
	}

	if !database.IsTaskStarted(db, task) {
		c.JSON(http.StatusConflict, Response{409, "error", "Task not started"})
		return
	}

	err = database.UpdateTask(db, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{500, "error", "Failed to stop task"})
		return
	}

	c.JSON(http.StatusOK, Response{200, "message", "Task stopped successfully"})
}
