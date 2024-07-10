package routers

import (
	"database/sql"
	"net/http"
	"strconv"
	"tz_iul/database"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func getUsers(c *gin.Context, db *sql.DB) {
	var users []database.User

	passportSerie := c.Query("passport_serie")
	passportNumber := c.Query("passport_number")
	surname := c.Query("surname")
	name := c.Query("name")
	patronymic := c.Query("patronymic")
	address := c.Query("address")

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

	query, err := db.Query(`
		SELECT id, surname, name, patronymic, address, passport_serie, passport_number
		FROM users
		WHERE (passport_serie = $1 OR $1 = '') AND
			(passport_number = $2 OR $2 = '') AND
			(surname = $3 OR $3 = '') AND
			(name = $4 OR $4 = '') AND
			(patronymic = $5 OR $5 = '') AND
			(address = $6 OR $6 = '')
		ORDER BY id ASC
		LIMIT $7 OFFSET $8;
	`, passportSerie, passportNumber, surname, name, patronymic, address, limit, (page-1)*limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	defer query.Close()

	var isFound bool
	for query.Next() {
		var user database.User
		err := query.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Patronymic, &user.Address, &user.PassportSerie, &user.PassportNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
			return
		}
		users = append(users, user)
		isFound = true
	}

	if !isFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	c.JSON(http.StatusOK, users)
}
