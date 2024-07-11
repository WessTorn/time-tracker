package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

// @Summary Get Tasks
// @Tags tasks
// @Description Retrieves a list of tasks for a user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} database.Task "List of tasks"
// @Failure 400 {object} Response "Invalid user ID"
// @Failure 404 {object} Response "Tasks not found"
// @Failure 500 {object} Response "Failed to get tasks"
// @Router /tasks [get]
func getTasks(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(getTasks)")

	getId := c.Param("id")

	id, err := strconv.Atoi(getId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid user ID"})
		return
	}

	tasks, err := database.SelectTasks(db, id)
	if err != nil {
		switch err.Error() {
		case "TasksNotFound":
			c.JSON(http.StatusNotFound, Response{404, "error", "Tasks not found"})
		default:
			c.JSON(http.StatusInternalServerError, Response{404, "error", "Failed to get tasks"})
		}
		return
	}

	c.JSON(http.StatusOK, tasks)
}
