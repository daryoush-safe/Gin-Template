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
				} else if registrationErrors, ok := err.(exceptions.UserRegistrationError); ok {
					handleRegistrationError(c, registrationErrors, recovery.constants.Translator)
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
	errorMessages := make(map[string]map[string]string)

	for _, validationError := range validationErrors {
		if _, ok := errorMessages[validationError.Field()]; !ok {
			errorMessages[validationError.Field()] = make(map[string]string)
		}
		errorMessages[validationError.Field()][validationError.Tag()] = validationError.Translate(trans)
	}

	controller.Response(c, 422, errorMessages, nil)
}

func handleBindingError(c *gin.Context, bindingError exceptions.BindingError, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	message, _ := trans.T("errors.generic")

	if numError, ok := bindingError.Err.(*strconv.NumError); ok {
		message, _ = trans.T("errors.numeric", numError.Num)
	}

	controller.Response(c, 400, message, nil)
}

func handleRegistrationError(c *gin.Context, registrationErrors exceptions.UserRegistrationError, transKey string) {
	trans := controller.GetTranslator(c, transKey)
	errorMessages := make(map[string]map[string]string)
	for _, registrationError := range registrationErrors.FieldErrors() {
		if _, ok := errorMessages[registrationError.Field]; !ok {
			errorMessages[registrationError.Field] = make(map[string]string)
		}
		message, _ := trans.T(registrationError.Tag) // I want to pass a field to json but I am dumb sadly!
		errorMessages[registrationError.Field][registrationError.Tag] = message
	}

	controller.Response(c, 422, errorMessages, nil)
}

func unhandledErrors(c *gin.Context, err error, transKey string) {
	log.Println(err.Error())
	trans := controller.GetTranslator(c, transKey)
	errorMessage, _ := trans.T("errors.generic")

	controller.Response(c, 500, errorMessage, nil)
}
