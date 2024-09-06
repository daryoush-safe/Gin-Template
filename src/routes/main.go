package routes

import (
	"first-project/src/middleware"
	routes_http_v1 "first-project/src/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.Use(middleware.Localization)
	ginEngine.Use(middleware.Recovery)

	v1 := ginEngine.Group("/v1")

	v1 = registerGeneralRoutes(v1)
	registerCustomerRoutes(v1)

	return ginEngine
}

func registerGeneralRoutes(v1 *gin.RouterGroup) *gin.RouterGroup {
	return routes_http_v1.SetupGeneralRoutes(v1)
}

func registerCustomerRoutes(v1 *gin.RouterGroup) *gin.RouterGroup {
	return routes_http_v1.SetupCustomerRoutes(v1.Group("customer"))
}
