package routes_http_v1

import (
	"github.com/gin-gonic/gin"

	pingPongController_v1 "first-project/first-project/controller/v1"
)

func SetupGneralRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("/ping", pingPongController_v1.Pong)

	return routerGroup
}
