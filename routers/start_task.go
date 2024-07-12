package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/WessTorn/time-tracker/database"
	"github.com/WessTorn/time-tracker/logger"

	"github.com/gin-gonic/gin"
)

// @Summary Start Task
// @Description Starts a task for a user
// @Tags tasks
// @Produce json
// @Param id path int true "User ID"
// @Param task_id body TaskID true "Task ID"
// @Success 200 {object} Response "Task started successfully"
// @Failure 400 {object} Response "Invalid user ID, Invalid request payload"
// @Failure 404 {object} Response "User not found"
// @Failure 409 {object} Response "Task already started"
// @Failure 500 {object} Response "Failed to start task"
// @Router /tasks/start/{id} [post]
func startTask(c *gin.Context, db *sql.DB) {
	logger.Log.Info("POST request (getUsers)")
	getId := c.Param("id")

	var request TaskID

	id, err := strconv.Atoi(getId)
	if err != nil {
		logger.Log.Debugf("(Atoi) %v", err)
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid user ID"})
		return
	}

	err = c.BindJSON(&request)
	if err != nil {
		logger.Log.Debugf("(BindJSON) %v", err)
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid request payload"})
		return
	}

	if !database.IsUserExistsFromID(db, id) {
		c.JSON(http.StatusNotFound, Response{404, "error", "User not found"})
		return
	}

	newTask := database.Task{
		UserID:    id,
		TaskID:    request.TaskID,
		StartTime: time.Now(),
	}

	if database.IsTaskStarted(db, newTask) {
		c.JSON(http.StatusConflict, Response{409, "error", "Task already started"})
		return
	}

	err = database.InsertTask(db, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{500, "error", "Failed to start task"})
		return
	}

	logger.Log.Debug("Reply to request: " + "Task started successfully")

	c.JSON(http.StatusOK, Response{200, "message", "Task started successfully"})
}
