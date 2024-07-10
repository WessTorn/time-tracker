package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func getUsers(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(getUsers)")

	var filterUser database.User

	filterUser.PassportSerie = c.Query("passport_serie")
	filterUser.PassportNumber = c.Query("passport_number")
	filterUser.LastName = c.Query("surname")
	filterUser.FirstName = c.Query("name")
	filterUser.Patronymic = c.Query("patronymic")
	filterUser.Address = c.Query("address")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid limit"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid page"})
		return
	}

	users, err := database.SelectUsers(db, filterUser, limit, page)
	if err != nil {
		switch err.Error() {
		case "UsersNotFound":
			c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		}
		return
	}

	c.JSON(http.StatusOK, users)
}
