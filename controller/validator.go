package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type customError struct {
// 	message interface{}
// }

// func (e *customError) Error() interface{} {
// 	return e.message
// }

func Validated[T any](c *gin.Context) T {
	var params T

	var validate *validator.Validate = validator.New()

	if err := c.ShouldBindUri(&params); err != nil {
		fmt.Println("khata")
	}

	if err := validate.Struct(params); err != nil {
		panic(err)
		// validationErrors := err.(validator.ValidationErrors)
		// for _, validationError := range validationErrors {
		//     fmt.Println(validationError.Error())
		// }
		// c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		// errorMsg := localization(params, err)

		// fmt.Println(&customError{message: errorMsg})
	}

	return params
}

// func localization[T any](params T, err error) interface{} {
// 	bundle := i18n.NewBundle(language.English)
// 	localizer := i18n.NewLocalizer(bundle, "en")
// 	translatedError := localizer.MustLocalize(&i18n.LocalizeConfig{
// 		DefaultMessage: &i18n.Message{
// 			ID:    "Erorr",
// 			Other: "$field should be numeric",
// 		},
// 	})

// 	var (
// 		g = galidator.New().CustomMessages(galidator.Messages{
// 			"numeric": translatedError,
// 		})
// 		customizer = g.Validator(params)
// 	)

// 	return customizer.DecryptErrors(err)
// }

// func isValid() {
// 	// Error()
// }
