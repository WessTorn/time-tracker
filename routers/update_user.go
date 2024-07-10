package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"tz_iul/database"

	"github.com/gin-gonic/gin"
)

func updateUser(c *gin.Context, db *sql.DB) {
	var user database.User
	id := c.Param("id")

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Формируем SQL запрос с учетом того, какие поля были изменены
	var sqlStatement string
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

	sqlStatement = fmt.Sprintf("UPDATE users SET %s WHERE id=%s;", strings.Join(fieldsToUpdate, ", "), id)
	result, err := db.Exec(sqlStatement, paramsToUpdate...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
