package validation

import (
	"first-project/src/bootstrap"
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func containsLowercase(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	matched, _ := regexp.MatchString("[a-z]", password)
	return matched
}

func containsUppercase(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	matched, _ := regexp.MatchString("[A-Z]", password)
	return matched
}

func containsNumber(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	matched, _ := regexp.MatchString("[0-9]", password)
	return matched
}

func containsSpecialChar(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	matched, _ := regexp.MatchString("[^\\d\\w]", password)
	return matched
}

func setCustomValidatorTranslator(validate *validator.Validate, trans ut.Translator, validatorKey string) {
	massage, _ := trans.T(validatorKey)
	validate.RegisterTranslation(validatorKey, trans, func(ut ut.Translator) error {
		return ut.Add(validatorKey, massage, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(validatorKey, fe.Field())
		return t
	})
}

func SetCustomValidators(validate *validator.Validate, constants *bootstrap.Context, trans ut.Translator) {
	validate.RegisterValidation(constants.ContainsLowercase, containsLowercase)
	setCustomValidatorTranslator(validate, trans, constants.ContainsLowercase)

	validate.RegisterValidation(constants.ContainsUppercase, containsUppercase)
	setCustomValidatorTranslator(validate, trans, constants.ContainsUppercase)

	validate.RegisterValidation(constants.ContainsNumber, containsNumber)
	setCustomValidatorTranslator(validate, trans, constants.ContainsNumber)

	validate.RegisterValidation(constants.ContainsSpecialChar, containsSpecialChar)
	setCustomValidatorTranslator(validate, trans, constants.ContainsSpecialChar)
}
