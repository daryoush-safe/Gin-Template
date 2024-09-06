package controller_v1_sample_general

import (
	application_math "first-project/src/application/math"
	"first-project/src/controller"

	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Add(c *gin.Context) {
	param := controller.Validated[AddParams](c)

	sum := application_math.Add(param.Num1, param.Num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
