package controller_v1_general

import (
	"first-project/src/application"
	application_communication "first-project/src/application/communication/emailService"
	"first-project/src/bootstrap"
	"first-project/src/controller"
	"first-project/src/exceptions"

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

func userRegistrationResponse(c *gin.Context, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("successMessage.userRegistration")
	controller.Response(c, 200, message, nil)
}

func emailVerificationResponse(c *gin.Context, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("successMessage.emailVerification")
	controller.Response(c, 200, message, nil)
}

func loginResponse(c *gin.Context, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("successMessage.login")
	controller.Response(c, 200, message, nil)
}

func forgotPasswordResponse(c *gin.Context, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("successMessage.forgotPassword")
	controller.Response(c, 200, message, nil)
}

func resetPasswordResponse(c *gin.Context, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("successMessage.resetPassword")
	controller.Response(c, 200, message, nil)
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

func (userController *UserController) sendResetPassEmail(c *gin.Context, email string) {
	templateData := struct {
		ResetLink string
	}{
		ResetLink: "http://localhost:8080/v1/resetPassword",
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
	userController.sendActivationEmail(c, param.Username, otp, param.Email)
	userController.userService.RegisterUser(param.Username, param.Email, param.Password, otp)
	userRegistrationResponse(c, userController.constants.Context.Translator)
}

func (userController *UserController) VerifyEmail(c *gin.Context) {
	type verifyEmailParams struct {
		OTP   string `json:"otp" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
	var registrationError exceptions.UserRegistrationError
	param := controller.Validated[verifyEmailParams](c, &userController.constants.Context)
	alreadyVerified := userController.userService.CheckUserAlreadyVerifiedByEmail(param.Email)
	if alreadyVerified {
		registrationError.AppendError(
			userController.constants.ErrorField.Email,
			userController.constants.ErrorTag.AlreadyVerified)
		panic(registrationError)
	}
	userController.otpService.VerifyOTP(param.OTP, param.Email)
	userController.userService.VerifyEmail(param.Email)
	emailVerificationResponse(c, userController.constants.Context.Translator)
}

func (userController *UserController) Login(c *gin.Context) {
	type loginParams struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	param := controller.Validated[loginParams](c, &userController.constants.Context)
	userController.userService.LoginService(param.Username, param.Password)
	loginResponse(c, userController.constants.Context.Translator)
}

func (userController *UserController) ForgotPassword(c *gin.Context) {
	type forgotPasswordParams struct {
		Email string `json:"email" validate:"required"`
	}
	param := controller.Validated[forgotPasswordParams](c, &userController.constants.Context)
	var registrationError exceptions.UserRegistrationError
	alreadyVerified := userController.userService.CheckUserAlreadyVerifiedByEmail(param.Email)
	if !alreadyVerified {
		registrationError.AppendError(
			userController.constants.ErrorField.Email,
			userController.constants.ErrorTag.EmailNotExist)
		panic(registrationError)
	}
	userController.sendResetPassEmail(c, param.Email)
	forgotPasswordResponse(c, userController.constants.Context.Translator)
}

func (userController *UserController) ResetPassword(c *gin.Context) {
	type resetPasswordParams struct {
		Email           string `json:"email" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirmPassword" validate:"required"`
	}
	param := controller.Validated[resetPasswordParams](c, &userController.constants.Context)
	userController.userService.ResetPasswordService(param.Email, param.Password, param.ConfirmPassword)
	resetPasswordResponse(c, userController.constants.Context.Translator)
}
