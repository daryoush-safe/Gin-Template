package controller_v1_general

import (
	"first-project/src/application"
	"first-project/src/bootstrap"
	"first-project/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	constants    *bootstrap.Constants
	userService  *application.UserService
	jwtService   *application.JwtService
	emailService *application.EmailService
}

func NewUserController(
	constants *bootstrap.Constants,
	userService *application.UserService,
	jwtService *application.JwtService,
	emailService *application.EmailService,
) *UserController {
	return &UserController{
		constants:    constants,
		userService:  userService,
		jwtService:   jwtService,
		emailService: emailService,
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
	// TODO: incorrect: remove jwt use otp instead
	tokenString := userController.jwtService.CreateToken(param.Email)
	userController.emailService.SendVerificationEmail(param.Username, param.Email, tokenString)
	// TODO: standard response
	// TODO: translate
	c.String(http.StatusOK, "Please verify your Email to activate your account!")
}

func (userController *UserController) VerifyEmail(c *gin.Context) {
	type verifyEmailParams struct {
		Token string `uri:"token" validate:"required"`
	}
	param := controller.Validated[verifyEmailParams](c, &userController.constants.Context)
	// TODO: use otp
	email := userController.jwtService.VerifyToken(param.Token)
	userController.userService.VerifyEmail(email)
	// TODO: standard response
	// TODO: translate
	c.String(http.StatusOK, "Email verified!")
}
