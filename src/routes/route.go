package routes

import (
	"first-project/src/bootstrap"
	middleware_exceptions "first-project/src/middleware/exceptions"
	middleware_i18n "first-project/src/middleware/i18n"
	middleware_rate_limit "first-project/src/middleware/rateLimit"
	routes_http_v1 "first-project/src/routes/http/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(ginEngine *gin.Engine, di *bootstrap.Di, db *gorm.DB) {
	localizationMiddleware := middleware_i18n.NewLocalization(&di.Constants.Context)
	recoveryMiddleware := middleware_exceptions.NewRecovery(&di.Constants.Context)
	rateLimitMiddleware := middleware_rate_limit.NewRateLimit(5, 10)

	ginEngine.Use(localizationMiddleware.Localization)
	ginEngine.Use(recoveryMiddleware.Recovery)
	ginEngine.Use(rateLimitMiddleware.RateLimit)

	v1 := ginEngine.Group("/v1")

	registerGeneralRoutes(v1, di, db)
	registerCustomerRoutes(v1, di)
}

func registerGeneralRoutes(v1 *gin.RouterGroup, di *bootstrap.Di, db *gorm.DB) *gin.RouterGroup {
	return routes_http_v1.SetupGeneralRoutes(v1, di, db)
}

func registerCustomerRoutes(v1 *gin.RouterGroup, di *bootstrap.Di) *gin.RouterGroup {
	return routes_http_v1.SetupCustomerRoutes(v1.Group("customer"), di)
}
