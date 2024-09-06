package routes_http_v1

import (
	"github.com/gin-gonic/gin"

	pingPongController_v1 "first-project/src/controller/v1"
)

func SetupCustomerRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("/ping1", pingPongController_v1.Pong)
	routerGroup.GET("/ping2", pingPongController_v1.Pong)
	routerGroup.GET("/ping3", pingPongController_v1.Pong)

	return routerGroup
}
