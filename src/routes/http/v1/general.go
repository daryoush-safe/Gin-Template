package routes_http_v1

import (
	"github.com/gin-gonic/gin"

	application_math "gingool/src/application/math"
	"gingool/src/bootstrap"
	controller_v1_general "gingool/src/controller/v1/general"
)

func SetupGeneralRoutes(routerGroup *gin.RouterGroup, di *bootstrap.Di, addService *application_math.AddService) *gin.RouterGroup {
	sampleController := controller_v1_general.NewSampleController(di.Constants, addService)

	routerGroup.GET("/ping", controller_v1_general.Pong)
	routerGroup.GET("/add/:num1/:num2", sampleController.Add)

	return routerGroup
}
