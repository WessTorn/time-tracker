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

// @Summary Add a new user
// @Tags users
// @Description Add a new user to the database using passport series and number
// @Accept json
// @Produce json
// @Param passportNumber body Passport true "Passport number"
// @Success 200 {object} Response "User added successfully"
// @Failure 400 {object} Response "Invalid request payload, Invalid passport number"
// @Failure 409 {object} Response "User already exists"
// @Failure 500 {object} Response "Failed to fetch user data from external API, Failed to add user to the database"
// @Router /users [post]
func addUser(c *gin.Context, db *sql.DB) {
	logger.Log.Debug("(addUser)")

	var request Passport
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid request payload"})
		return
	}

	serie, number, check := checkPassportNumber(request.PassportNumber)

	if !check {
		c.JSON(http.StatusBadRequest, Response{400, "error", "Invalid passport number"})
		return
	}

	var getUser *database.User
	getUser, err = GetUserDataFromExternalAPI(serie, number)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{500, "error", "Failed to fetch user data from external API"})
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
		c.JSON(http.StatusConflict, Response{409, "error", "User already exists"})
		return
	}

	err = database.InsertUser(db, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{500, "error", "Failed to add user to the database"})
		return
	}

	c.JSON(http.StatusOK, Response{200, "message", "User added successfully"})
}

func checkPassportNumber(passportNumber string) (serie string, number string, check bool) {
	logger.Log.Debug("(checkPassportNumber)")
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
