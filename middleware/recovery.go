package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ExceptionHandler(c *gin.Context) {
	defer func() {
		// Catch panic (if any) and handle it as an error
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
	// Call the next handler in the chain
	c.Next()
}

func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors) {
	for _, validationError := range validationErrors {
		log.Println(validationError.Error())
		c.String(422, validationError.Error())
	}
}

func unhandledErrors(c *gin.Context, err error) {
	log.Println(err.Error())
	c.String(500, "error")
}
