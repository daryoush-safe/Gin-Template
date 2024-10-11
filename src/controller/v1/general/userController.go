package controller_v1_general

import (
	"first-project/src/application"
	"first-project/src/bootstrap"
	"first-project/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	constants   *bootstrap.Constants
	userService *application.UserService
}

func NewUserController(constants *bootstrap.Constants, userService *application.UserService) *UserController {
	return &UserController{
		constants:   constants,
		userService: userService,
	}
}

func (userController *UserController) Register(c *gin.Context) {
	type registerParams struct {
		Username string `json:"username" validate:"required,gt=2,lt=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	param := controller.Validated[registerParams](c, &userController.constants.Context)
	userController.userService.RegisterService(param.Username, param.Email, param.Password)
	c.String(http.StatusOK, "Registered Successfully!")
}
