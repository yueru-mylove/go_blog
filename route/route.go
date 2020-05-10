package route

import (
	"awesomeProject/package/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(settings.RunMode)
	r.GET("test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message" : "test",})
	})

	return r
}