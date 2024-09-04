package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golodash/galidator"
)

type AddParams struct {
	Num1 string `uri:"num1" validate:"required,numeric"`
	Num2 string `uri:"num2" validate:"required,numeric"`
}

func ValidateNunmber(c *gin.Context) bool {
	var (
		g = galidator.New().CustomMessages(galidator.Messages{
			"numeric": "$field should be numeric",
		})
		customizer = g.Validator(AddParams{})
	)
	var validate *validator.Validate = validator.New()
	var params AddParams

	if err := c.ShouldBindUri(&params); err != nil {
		return false
	}

	if err := validate.Struct(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customizer.DecryptErrors(err)})
		return false
	}

	return true
}
