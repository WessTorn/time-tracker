package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

func updateUser(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(updateUser)")
	id := c.Param("id")

	var user database.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	setMessage := strings.Join(fieldsToUpdate, ", ")

	err = database.UpdateUser(db, id, setMessage, paramsToUpdate)
	if err != nil {
		switch err.Error() {
		case "UserNotFound":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
