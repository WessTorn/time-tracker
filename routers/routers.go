package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time-tracker/config"
	"time-tracker/database"
	"time-tracker/logger"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) *gin.Engine {
	logger.Log.Debug("(InitRouter)")

	router := gin.Default()

	router.GET(config.HostRelativePath(), func(c *gin.Context) {
		getUsers(c, db) // 1) Получение данных пользователей
	})

	router.GET(config.HostRelativePath()+"/task/:id", func(c *gin.Context) {
		getTasks(c, db) // 2) Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей
	})

	router.POST(config.HostRelativePath()+"/task/start/:id", func(c *gin.Context) {
		startTask(c, db) // 3) Начать отсчет времени по задаче для пользователя
	})

	router.POST(config.HostRelativePath()+"/task/stop/:id", func(c *gin.Context) {
		stopTask(c, db) // 4) Закончить отсчет времени по задаче для пользователя
	})

	router.DELETE(config.HostRelativePath()+"/:id", func(c *gin.Context) {
		deleteUser(c, db) // 5) Удаление пользователя
	})

	router.PUT(config.HostRelativePath()+"/:id", func(c *gin.Context) {
		updateUser(c, db) // 6) Изменение данных пользователя
	})

	router.POST(config.HostRelativePath(), func(c *gin.Context) {
		addUser(c, db) // 7) Добавление нового пользователя
	})

	return router
}

func GetUserDataFromExternalAPI(serie string, number string) (*database.User, error) {
	logger.Log.Debug("(GetUserDataFromExternalAPI)")

	url := config.ExternalApiURL() + "?passportSerie=" + serie + "&passportNumber=" + number

	logger.Log.Debugf("(url) %s", url)

	resp, err := http.Get(url)

	if err != nil {
		logger.Log.Debugf("(Get) %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Log.Debugf("(Status) %d", resp.StatusCode)
		return nil, err
	}

	var userData *database.User
	err = json.NewDecoder(resp.Body).Decode(&userData)

	fmt.Println(userData.FirstName, userData.LastName, userData.Patronymic, userData.Address)

	if err != nil {
		logger.Log.Debugf("(Decode) %v", err)
		return nil, err
	}

	return userData, nil
}
