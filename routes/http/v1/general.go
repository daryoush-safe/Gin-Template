package routes_http_v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupGneralRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return routerGroup
}
