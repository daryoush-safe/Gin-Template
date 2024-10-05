package bootstrap

import (
	"fmt"
)

type Constants struct {
	Context Context
	Redis   Redis
}

type Context struct {
	Translator                    string
	IsLoadedValidationTranslator  string
	IsLoadedCustomValidationError string
	ContainsLowercase             string
	ContainsUppercase             string
	ContainsNumber                string
	ContainsSpecialChar           string
}

type Redis struct {
}

func NewConstants() *Constants {
	return &Constants{
		Context: Context{
			Translator:                    "translator",
			IsLoadedValidationTranslator:  "isLoadedValidationTranslator",
			IsLoadedCustomValidationError: "isLoadedCustomValidationError",
			ContainsLowercase:             "containsLowercase",
			ContainsUppercase:             "containsUppercase",
			ContainsNumber:                "containsNumber",
			ContainsSpecialChar:           "containsSpecialChar",
		},
		Redis: Redis{},
	}
}

func (r *Redis) GetUserID(userID int) string {
	return fmt.Sprintf("user:%d", userID)
}
