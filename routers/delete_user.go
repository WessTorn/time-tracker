package routers

import (
	"database/sql"
	"net/http"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

func deleteUser(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("deleteUser")
	id := c.Param("id")

	err := database.DeletUserFromID(db, id)

	if err != nil {
		switch err.Error() {
		case "UserNotFound":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
