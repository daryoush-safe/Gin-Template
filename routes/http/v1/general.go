package routes_http_v1

import (
	"github.com/gin-gonic/gin"

	controller_v1 "first-project/controller/v1"
)

func SetupGneralRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("/ping", controller_v1.Pong)
	routerGroup.GET("/add/:num1/:num2", controller_v1.Add)

	return routerGroup
}
