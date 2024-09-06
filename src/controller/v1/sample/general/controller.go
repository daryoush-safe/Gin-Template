package controller_v1_sample_general

import (
	application_math "first-project/src/application/math"
	"first-project/src/controller"

	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Add(c *gin.Context) {
	controller.Validated[AddParams](c)

	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(c.Param("num2"))

	sum := application_math.Add(num1, num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
