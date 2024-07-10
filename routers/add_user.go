package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

func addUser(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(addUser)")

	var request struct {
		PassportNumber string `json:"passportNumber"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	serie, number, check := checkPassportNumber(request.PassportNumber)

	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid passport number"})
		return
	}

	var getUser *database.User
	getUser, err = GetUserDataFromExternalAPI(serie, number)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data from external API"})
		return
	}

	newUser := database.User{
		LastName:       getUser.LastName,
		FirstName:      getUser.FirstName,
		Patronymic:     getUser.Patronymic,
		Address:        getUser.Address,
		PassportSerie:  serie,
		PassportNumber: number,
	}

	if database.IsUserExists(db, newUser) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
		return
	}

	err = database.InsertUser(db, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})

}

func checkPassportNumber(passportNumber string) (serie string, number string, check bool) {
	parts := strings.Split(passportNumber, " ")
	if len(parts) != 2 {
		return "", "", false
	}

	if len(parts[0]) != 4 || len(parts[1]) != 6 {
		return "", "", false
	}

	_, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", "", false
	}

	_, err = strconv.Atoi(parts[1])
	if err != nil {
		return "", "", false
	}

	return parts[0], parts[1], true
}
