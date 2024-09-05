package routes

import (
	"first-project/middleware"
	routes_http_v1 "first-project/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	ginEngine := gin.Default()

	ginEngine.Use(middleware.ExceptionHandler)
	v1 := ginEngine.Group("/v1")

	v1 = registerGeneralRoutes(v1)
	registerCustomerRoutes(v1)

	return ginEngine
}

func registerGeneralRoutes(v1 *gin.RouterGroup) *gin.RouterGroup {
	return routes_http_v1.SetupGneralRoutes(v1)
}

func registerCustomerRoutes(v1 *gin.RouterGroup) *gin.RouterGroup {
	return routes_http_v1.SetupCustomerRoutes(v1.Group("customer"))
}
