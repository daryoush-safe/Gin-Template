package routes

import (
	middleware_exceptions "first-project/src/middleware/exceptions"
	middleware_i18n "first-project/src/middleware/i18n"
	routes_http_v1 "first-project/src/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.Use(middleware_i18n.Localization)
	ginEngine.Use(middleware_exceptions.Recovery)

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
