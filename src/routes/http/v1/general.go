package routes_http_v1

import (
	"github.com/gin-gonic/gin"

	controller_v1 "first-project/src/controller/v1"
	controller_v1_sample_general "first-project/src/controller/v1/sample/general"
)

func SetupGeneralRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("/ping", controller_v1.Pong)
	routerGroup.GET("/add/:num1/:num2", controller_v1_sample_general.Add)

	return routerGroup
}
