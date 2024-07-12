package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/WessTorn/time-tracker/database"
	"github.com/WessTorn/time-tracker/logger"

	"github.com/gin-gonic/gin"
)

// @Summary Update User
// @Tags users
// @Description Updates a user by ID
// @Produce json
// @Param id path string true "User ID"
// @Param user body database.User true "User data to update"
// @Success 200 {object} Response "User updated successfully"
// @Failure 400 {object} Response "No fields to update, Error print"
// @Failure 404 {object} Response "User not found"
// @Failure 500 {object} Response "Failed to update user"
// @Router /users/{id} [put]
func updateUser(c *gin.Context, db *sql.DB) {
	logger.Log.Info("PUT request (updateUser)")
	id := c.Param("id")

	var user database.User

	err := c.BindJSON(&user)
	if err != nil {
		logger.Log.Debugf("(BindJSON) %v", err)
		c.JSON(http.StatusBadRequest, Response{400, "error", err.Error()})
		return
	}

	fieldsToUpdate := make([]string, 0, 6)
	paramsToUpdate := make([]interface{}, 0, 6)

	var i int = 0

	if user.LastName != "" {
		i++
		value := fmt.Sprintf("surname=$%d", i)
		fieldsToUpdate = append(fieldsToUpdate, value)
		paramsToUpdate = append(paramsToUpdate, user.LastName)
	}

	if user.FirstName != "" {
		i++
		value := fmt.Sprintf("name=$%d", i)
		fieldsToUpdate = append(fieldsToUpdate, value)
		paramsToUpdate = append(paramsToUpdate, user.FirstName)
	}

	if user.Patronymic != "" {
		i++
		value := fmt.Sprintf("patronymic=$%d", i)
		fieldsToUpdate = append(fieldsToUpdate, value)
		paramsToUpdate = append(paramsToUpdate, user.Patronymic)
	}

	if user.Address != "" {
		i++
		value := fmt.Sprintf("address=$%d", i)
		fieldsToUpdate = append(fieldsToUpdate, value)
		paramsToUpdate = append(paramsToUpdate, user.Address)
	}

	if user.PassportSerie != "" {
		i++
		value := fmt.Sprintf("passport_serie=$%d", i)
		fieldsToUpdate = append(fieldsToUpdate, value)
		paramsToUpdate = append(paramsToUpdate, user.PassportSerie)
	}

	if user.PassportNumber != "" {
		i++
		value := fmt.Sprintf("passport_number=$%d", i)
		fieldsToUpdate = append(fieldsToUpdate, value)
		paramsToUpdate = append(paramsToUpdate, user.PassportNumber)
	}

	if len(fieldsToUpdate) == 0 {
		logger.Log.Debugf("(updateUser) No fields to update")
		c.JSON(http.StatusBadRequest, Response{400, "error", "No fields to update"})
		return
	}

	setMessage := strings.Join(fieldsToUpdate, ", ")

	err = database.UpdateUser(db, id, setMessage, paramsToUpdate)
	if err != nil {
		switch err.Error() {
		case "UserNotFound":
			c.JSON(http.StatusNotFound, Response{404, "error", "User not found"})
		default:
			c.JSON(http.StatusInternalServerError, Response{500, "error", "Failed to update user"})
			return
		}
		return
	}

	logger.Log.Debugf("Reply to request: User updated successfully")

	c.JSON(http.StatusOK, Response{200, "message", "User updated successfully"})
}
