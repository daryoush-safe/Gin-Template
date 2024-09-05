package controller

import (
	"first-project/localization"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/fa"
)

func Validated[T any](c *gin.Context) T {

	var params T
	var validate *validator.Validate = validator.New()

	trans := localization.GetTranslator()
	if trans.Locale() == "fa_IR" {
		fa.RegisterDefaultTranslations(validate, trans)
	} else {
		en.RegisterDefaultTranslations(validate, trans)
	}

	if err := c.ShouldBindUri(&params); err != nil {
		fmt.Println("error")
	}

	if err := validate.Struct(params); err != nil {
		panic(err)
	}

	return params
}
