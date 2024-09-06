package middleware

import (
	"first-project/src/localization"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ExceptionHandler(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {
					handleValidationError(c, validationErrors)
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

	c.JSON(422, validationErrors.Translate(trans))
}

func unhandledErrors(c *gin.Context, err error) {
	log.Println(err.Error())
	trans := localization.GetTranslator()
	errorMessage, _ := trans.T("errors.generic")
	c.String(500, errorMessage)
}
