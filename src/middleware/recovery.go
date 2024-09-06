package middleware

import (
	"first-project/src/controller"
	"first-project/src/exceptions"
	"first-project/src/localization"
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
	trans := localization.GetTranslator()

	c.JSON(422, controller.ValidationResponse{
		Status:   422,
		Messages: validationErrors.Translate(trans),
	})
}

func handleBindingError(c *gin.Context, bindingError exceptions.BindingError) {
	trans := localization.GetTranslator()
	message, _ := trans.T("errors.generic")

	if numError, ok := bindingError.Err.(*strconv.NumError); ok {
		message, _ = trans.T("errors.numeric", numError.Num)
	}

	c.JSON(400, controller.Response{
		Status:  400,
		Message: message,
		Data:    nil,
	})
}

func unhandledErrors(c *gin.Context, err error) {
	log.Println(err.Error())
	trans := localization.GetTranslator()
	errorMessage, _ := trans.T("errors.generic")

	c.JSON(500, controller.Response{
		Status:  500,
		Message: errorMessage,
		Data:    nil,
	})
}
