package controller_v1_general

import (
	application_math "first-project/src/application/math"
	"first-project/src/bootstrap"
	"first-project/src/controller"

	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

type SampleController struct {
	Constants *bootstrap.Constants
}

func (sample *SampleController) Add(c *gin.Context) {
	type AddParams struct {
		Num1 int `uri:"num1" validate:"required,number,gt=10"`
		Num2 int `uri:"num2" validate:"required,number"`
	}
	param := controller.Validated[AddParams](c)

	sum := application_math.Add(param.Num1, param.Num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
