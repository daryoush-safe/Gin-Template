package middleware_i18n

import (
	"first-project/src/localization"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

func Localization(c *gin.Context) {
	locale := getLocale(c.Request)
	c.Set("translator", localization.GetTranslator(locale))

	c.Next()
}

func getLocale(request *http.Request) string {
	return request.Header.Get("Accept-Language")
}

func GetTranslator(c *gin.Context) ut.Translator {
	translator, exists := c.Get("translator")
	if !exists {
		panic("translator not registered!")
	}

	return translator.(ut.Translator)
}
