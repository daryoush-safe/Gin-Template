package middleware_exceptions

import (
	"first-project/src/bootstrap"
	"first-project/src/controller"
	"first-project/src/exceptions"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RecoveryMiddleware struct {
	constants *bootstrap.Context
}

func NewRecovery(constants *bootstrap.Context) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		constants: constants,
	}
}

func (recovery RecoveryMiddleware) Recovery(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {
					handleValidationError(c, validationErrors, recovery.constants.Translator)
				} else if bindingError, ok := err.(exceptions.BindingError); ok {
					handleBindingError(c, bindingError, recovery.constants.Translator)
				} else if registrationError, ok := err.(exceptions.UserRegistrationError); ok {
					handleRegistrationError(c, registrationError, recovery.constants.Translator)
				} else {
					unhandledErrors(c, err, recovery.constants.Translator)
				}

				c.Abort()
			}
		}
	}()

	c.Next()
}

func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	errorsMessages := make(map[string]string)

	for _, validationError := range validationErrors {
		errorsMessages[validationError.Field()] = validationError.Translate(trans)
	}

	controller.Response(c, 422, errorsMessages, nil)
}

func handleBindingError(c *gin.Context, bindingError exceptions.BindingError, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("errors.generic")

	if numError, ok := bindingError.Err.(*strconv.NumError); ok {
		message, _ = trans.T("errors.numeric", numError.Num)
	}

	controller.Response(c, 400, message, nil)
}

func handleRegistrationError(c *gin.Context, registrationError exceptions.UserRegistrationError, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	errorsMessages := make(map[string]string)
	if registrationError.Username != "" {
		message, _ := trans.T("user_registration.username_exists")
		errorsMessages["username"] = message
	}
	if registrationError.Email != "" {
		message, _ := trans.T("user_registration.email_exists")
		errorsMessages["email"] = message
	}
	controller.Response(c, 422, errorsMessages, nil)
}

func unhandledErrors(c *gin.Context, err error, transKey string) {
	log.Println(err.Error())
	trans := controller.GetTranslator(c, transKey)
	errorMessage, _ := trans.T("errors.generic")

	controller.Response(c, 500, errorMessage, nil)
}
