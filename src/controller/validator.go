package controller

import (
	"first-project/src/exceptions"
	middleware_i18n "first-project/src/middleware/i18n"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/fa"
)

var validate *validator.Validate = validator.New()

func setupTranslation(c *gin.Context) {
	trans := middleware_i18n.GetTranslator(c)
	key := "isLoadedValidationTranslator"

	_, exists := c.Get(key)
	if !exists {
		if trans.Locale() == "fa_IR" {
			fa.RegisterDefaultTranslations(validate, trans)
		} else {
			en.RegisterDefaultTranslations(validate, trans)
		}

		c.Set(key, true)
	}
}

func Validated[T any](c *gin.Context) T {
	setupTranslation(c)

	var params T
	if err := c.ShouldBindUri(&params); err != nil {
		bindingError := exceptions.BindingError{Err: err}
		panic(bindingError)
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		bindingError := exceptions.BindingError{Err: err}
		panic(bindingError)
	}

	if err := c.ShouldBind(&params); err != nil {
		bindingError := exceptions.BindingError{Err: err}
		panic(bindingError)
	}

	if err := validate.Struct(params); err != nil {
		panic(err)
	}

	return params
}
