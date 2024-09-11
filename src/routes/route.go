package routes

import (
	"gingool/src/bootstrap"
	middleware_exceptions "gingool/src/middleware/exceptions"
	middleware_i18n "gingool/src/middleware/i18n"
	routes_http_v1 "gingool/src/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run(ginEngine *gin.Engine, di *bootstrap.Di) {
	localizationMiddleware := middleware_i18n.NewLocalization(&di.Constants.Context)
	recoveryMiddleware := middleware_exceptions.NewRecovery(&di.Constants.Context)

	ginEngine.Use(localizationMiddleware.Localization)
	ginEngine.Use(recoveryMiddleware.Recovery)

	v1 := ginEngine.Group("/v1")

	registerGeneralRoutes(v1, di)
	registerCustomerRoutes(v1, di)
}

func registerGeneralRoutes(v1 *gin.RouterGroup, di *bootstrap.Di) *gin.RouterGroup {
	return routes_http_v1.SetupGeneralRoutes(v1, di)
}

func registerCustomerRoutes(v1 *gin.RouterGroup, di *bootstrap.Di) *gin.RouterGroup {
	return routes_http_v1.SetupCustomerRoutes(v1.Group("customer"), di)
}
