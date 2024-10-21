package controller_v1_general

import (
	"first-project/src/application"
	application_communication "first-project/src/application/communication/emailService"
	"first-project/src/bootstrap"
	"first-project/src/controller"
	"first-project/src/jwt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	constants    *bootstrap.Constants
	userService  *application.UserService
	otpService   *application.OTPService
	emailService *application_communication.EmailService
}

func NewUserController(
	constants *bootstrap.Constants,
	userService *application.UserService,
	otpService *application.OTPService,
	emailService *application_communication.EmailService,
) *UserController {
	return &UserController{
		constants:    constants,
		userService:  userService,
		otpService:   otpService,
		emailService: emailService,
	}
}

func setupResponse(c *gin.Context, transKey string, messageTag string, data interface{}) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T(messageTag)
	controller.Response(c, 200, message, data)
}

func getTemplatePath(c *gin.Context, transKey string) string {
	trans := controller.GetTranslator(c, transKey)
	if trans.Locale() == "fa_IR" {
		return "fa.html"
	}
	return "en.html"
}

func (userController *UserController) sendActivationEmail(c *gin.Context, username, otp, email string) {
	templateData := struct {
		Username       string
		OTP            string
		ActivationLink string
	}{
		Username:       username,
		OTP:            otp,
		ActivationLink: "http://localhost:8080/v1/register/activate",
	}
	templatePath := getTemplatePath(c, userController.constants.Context.Translator)
	// you can also add constant value for subjects!
	userController.emailService.SendEmail(
		email, "Activate account", "activateAccount/"+templatePath, templateData)
}

func (userController *UserController) sendResetPassEmail(c *gin.Context, email, token string) {
	templateData := struct {
		ResetLink string
	}{
		ResetLink: "http://localhost:8080/v1/resetPassword?token=" + token,
	}
	templatePath := getTemplatePath(c, userController.constants.Context.Translator)
	// you can also add constant value for subjects!
	userController.emailService.SendEmail(
		email, "Reset Password", "resetPassword/"+templatePath, templateData)
}

func (userController *UserController) Register(c *gin.Context) {
	type registerParams struct {
		Username        string `json:"username" validate:"required,gt=2,lt=20"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirmPassword" validate:"required"`
	}
	param := controller.Validated[registerParams](c, &userController.constants.Context)
	userController.userService.VerifyUserRegistration(param.Username, param.Email, param.Password, param.ConfirmPassword)
	otp := application.GenerateOTP()
	userController.userService.RegisterUser(param.Username, param.Email, param.Password, otp)
	userController.sendActivationEmail(c, param.Username, otp, param.Email)
	setupResponse(c, userController.constants.Context.Translator, "successMessage.userRegistration", nil)
}

func (userController *UserController) VerifyEmail(c *gin.Context) {
	type verifyEmailParams struct {
		OTP   string `json:"otp" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
	param := controller.Validated[verifyEmailParams](c, &userController.constants.Context)
	userController.userService.VerifyUserNotExist(param.Email)
	userController.otpService.VerifyOTP(param.OTP, param.Email)
	userController.userService.VerifyEmail(param.Email)
	setupResponse(c, userController.constants.Context.Translator, "successMessage.emailVerification", nil)
}

func (userController *UserController) Login(c *gin.Context) {
	type loginParams struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	param := controller.Validated[loginParams](c, &userController.constants.Context)
	userController.userService.LoginService(param.Username, param.Password)
	jwtString := jwt.GenerateJWT(c, "./jwtKeys", userController.constants.Context.IsLoadedJWTPrivateKey, param.Username)
	setupResponse(c, userController.constants.Context.Translator, "successMessage.login", jwtString)
}

func (userController *UserController) ForgotPassword(c *gin.Context) {
	type forgotPasswordParams struct {
		Email string `json:"email" validate:"required,email"`
	}
	param := controller.Validated[forgotPasswordParams](c, &userController.constants.Context)
	token := application.GenerateOTP()
	userController.userService.ForgotPasswordService(param.Email, token)
	userController.sendResetPassEmail(c, param.Email, token)
	setupResponse(c, userController.constants.Context.Translator, "successMessage.forgotPassword", nil)
}

func (userController *UserController) ResetPassword(c *gin.Context) {
	type resetPasswordParams struct {
		Email           string `json:"email" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirmPassword" validate:"required"`
		Token           string `form:"token" validate:"required"`
	}
	param := controller.Validated[resetPasswordParams](c, &userController.constants.Context)
	userController.otpService.VerifyOTP(param.Token, param.Email)
	userController.userService.ResetPasswordService(param.Email, param.Password, param.ConfirmPassword, param.Token)
	setupResponse(c, userController.constants.Context.Translator, "successMessage.resetPassword", nil)
}
