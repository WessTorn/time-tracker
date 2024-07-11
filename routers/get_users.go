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

// @Summary Get Users
// @Tags users
// @Description Retrieves a list of users with optional filters
// @Produce json
// @Param passport_serie query string false "Passport Series"
// @Param passport_number query string false "Passport Number"
// @Param surname query string false "Last Name"
// @Param name query string false "First Name"
// @Param patronymic query string false "Patronymic"
// @Param address query string false "Address"
// @Param limit query int false "Limit" default 10
// @Param page query int false "Page" default 1
// @Success 200 {object} database.User "List of users"
// @Failure 400 {object} Response "Invalid limit, Invalid page"
// @Failure 404 {object} Response "Users not found"
// @Failure 500 {object} Response "Failed to get users"
// @Router /users [get]
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
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid limit"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid page"})
		return
	}

	users, err := database.SelectUsers(db, filterUser, limit, page)
	if err != nil {
		switch err.Error() {
		case "UsersNotFound":
			c.JSON(http.StatusNotFound, Response{404, "error", "Users not found"})
		default:
			c.JSON(http.StatusInternalServerError, Response{404, "error", "Failed to get users"})
		}
		return
	}

	c.JSON(http.StatusOK, users)
}
