package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"tz_iul/database"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		getUsers(c, db)
	})

	router.POST("/users", func(c *gin.Context) {
		addUser(c, db)
	})

	return router
}

func GetUserDataFromExternalAPI(serie string, number string) (*database.User, error) {
	url := "http://localhost:8081/info" + "?passportSerie=" + serie + "&passportNumber=" + number
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var userData *database.User
	err = json.NewDecoder(resp.Body).Decode(&userData)

	fmt.Println(userData.FirstName, userData.LastName, userData.Patronymic, userData.Address)

	if err != nil {
		return nil, err
	}

	return userData, nil
}
