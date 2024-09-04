package routes_http_v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("ping1", func(c *gin.Context) {
		c.String(http.StatusOK, "pong1")
	})
	routerGroup.GET("ping2", func(c *gin.Context) {
		c.String(http.StatusOK, "pong2")
	})
	routerGroup.GET("ping3", func(c *gin.Context) {
		c.String(http.StatusOK, "pong3")
	})

	return routerGroup
}
