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
	otpService   *application.OTPService
	emailService *application.EmailService
}

func NewUserController(
	constants *bootstrap.Constants,
	userService *application.UserService,
	otpService *application.OTPService,
	emailService *application.EmailService,
) *UserController {
	return &UserController{
		constants:    constants,
		userService:  userService,
		otpService:   otpService,
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
	userController.userService.VerifyUserRegistration(param.Username, param.Email, param.Password)
	otp := application.GenerateOTP()
	userController.emailService.SendVerificationEmail(param.Username, param.Email, otp)
	userController.userService.RegisterUser(param.Username, param.Email, param.Password, otp)
	// TODO: standard response
	// TODO: translate
	c.String(http.StatusOK, "Please verify your Email to activate your account!")
}

func (userController *UserController) VerifyEmail(c *gin.Context) {
	type verifyEmailParams struct {
		OTP   string `json:"otp" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
	param := controller.Validated[verifyEmailParams](c, &userController.constants.Context)
	userController.userService.CheckUserAlreadyVerified(param.Email)
	userController.otpService.VerifyOTP(param.OTP, param.Email)
	userController.userService.VerifyEmail(param.Email)
	// TODO: standard response
	// TODO: translate
	c.String(http.StatusOK, "Email verified!")
}
