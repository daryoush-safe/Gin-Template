package routes

import (
	"fmt"
	application_math "gingool/src/application/math"
	"gingool/src/bootstrap"
	middleware_exceptions "gingool/src/middleware/exceptions"
	middleware_i18n "gingool/src/middleware/i18n"
	"gingool/src/repository"
	routes_http_v1 "gingool/src/routes/http/v1"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run(ginEngine *gin.Engine, di *bootstrap.Di) {
	localizationMiddleware := middleware_i18n.NewLocalization(&di.Constants.Context)
	recoveryMiddleware := middleware_exceptions.NewRecovery(&di.Constants.Context)

	ginEngine.Use(localizationMiddleware.Localization)
	ginEngine.Use(recoveryMiddleware.Recovery)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		di.Env.PRIMARY_DB.DB_USER,
		di.Env.PRIMARY_DB.DB_PASS,
		di.Env.PRIMARY_DB.DB_HOST,
		di.Env.PRIMARY_DB.DB_PORT,
		di.Env.PRIMARY_DB.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if(err != nil) {
		log.Fatal("Application cannot connect to database")
	}

	userRepository := repository.NewUserRepository(db)
	addService := application_math.NewAddService(userRepository)

	v1 := ginEngine.Group("/v1")

	registerGeneralRoutes(v1, di, addService)
	registerCustomerRoutes(v1, di)
}

func registerGeneralRoutes(v1 *gin.RouterGroup, di *bootstrap.Di, addService *application_math.AddService) *gin.RouterGroup {
	return routes_http_v1.SetupGeneralRoutes(v1, di, addService)
}

func registerCustomerRoutes(v1 *gin.RouterGroup, di *bootstrap.Di) *gin.RouterGroup {
	return routes_http_v1.SetupCustomerRoutes(v1.Group("customer"), di)
}
