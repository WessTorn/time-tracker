package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {

	})

	return router
}
