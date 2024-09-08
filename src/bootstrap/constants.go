package bootstrap

import (
	"fmt"
)

type Constants struct {
	Context Context
	Redis   Redis
}

type Context struct {
	Translator                   string
	IsLoadedValidationTranslator string
}

type Redis struct {
}

func NewConstants() *Constants {
	return &Constants{
		Context: Context{
			Translator:                   "translator",
			IsLoadedValidationTranslator: "isLoadedValidationTranslator",
		},
		Redis: Redis{},
	}
}

func (r *Redis) GetUserID(userID int) string {
	return fmt.Sprintf("user:%d", userID)
}
