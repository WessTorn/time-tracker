package routers

import (
	"database/sql"
	"net/http"

	"github.com/WessTorn/time-tracker/database"
	"github.com/WessTorn/time-tracker/logger"

	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Tags users
// @Description Deletes a user by ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response "User deleted successfully"
// @Failure 404 {object} Response "User not found"
// @Failure 500 {object} Response "Failed to delete user"
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context, db *sql.DB) {
	logger.Log.Info("DELETE request (deleteUser)")
	id := c.Param("id")

	err := database.DeletUserFromID(db, id)

	if err != nil {
		switch err.Error() {
		case "UserNotFound":
			c.JSON(http.StatusNotFound, Response{404, "error", "User not found"})
		default:
			c.JSON(http.StatusInternalServerError, Response{500, "error", "Failed to delete user"})
		}
		return
	}

	logger.Log.Debug("Reply to request: " + "User deleted successfully")

	c.JSON(http.StatusOK, Response{200, "message", "User deleted successfully"})
}
