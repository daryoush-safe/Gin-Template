package middleware_exceptions

import (
	"first-project/src/controller"
	"first-project/src/exceptions"
	middleware_i18n "first-project/src/middleware/i18n"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recovery(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {
					handleValidationError(c, validationErrors)
				} else if bindingError, ok := err.(exceptions.BindingError); ok {
					handleBindingError(c, bindingError)
				} else {
					unhandledErrors(c, err)
				}
			}
		}
	}()

	c.Next()
}

func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors) {
	trans := middleware_i18n.GetTranslator(c)
	errorsMessages := make(map[string]string)

	for _, validationError := range validationErrors {
		errorsMessages[validationError.Field()] = validationError.Translate(trans)
	}

	controller.Response(c, 422, errorsMessages, nil)
}

func handleBindingError(c *gin.Context, bindingError exceptions.BindingError) {
	trans := middleware_i18n.GetTranslator(c)
	message, _ := trans.T("errors.generic")

	if numError, ok := bindingError.Err.(*strconv.NumError); ok {
		message, _ = trans.T("errors.numeric", numError.Num)
	}

	controller.Response(c, 400, message, nil)
}

func unhandledErrors(c *gin.Context, err error) {
	log.Println(err.Error())
	trans := middleware_i18n.GetTranslator(c)
	errorMessage, _ := trans.T("errors.generic")

	controller.Response(c, 500, errorMessage, nil)
}
