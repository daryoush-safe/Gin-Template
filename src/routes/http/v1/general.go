package routes_http_v1

import (
	"github.com/gin-gonic/gin"

	"first-project/src/bootstrap"
	controller_v1_general "first-project/src/controller/v1/general"
)

func SetupGeneralRoutes(routerGroup *gin.RouterGroup, di *bootstrap.Di) *gin.RouterGroup {
	sampleController := &controller_v1_general.SampleController{
		Constants: di.Constants,
	}

	routerGroup.GET("/ping", controller_v1_general.Pong)
	routerGroup.GET("/add/:num1/:num2", sampleController.Add)

	return routerGroup
}
