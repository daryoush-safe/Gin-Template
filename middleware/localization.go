package middleware

import (
	"first-project/localization"

	"github.com/gin-gonic/gin"
)

func Localization(c *gin.Context) {
	localization.Register(c.Request)

	c.Next()
}
