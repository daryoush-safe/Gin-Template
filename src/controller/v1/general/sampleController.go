package controller_v1_general

import (
	application_math "first-project/src/application/math"
	"first-project/src/bootstrap"
	"first-project/src/controller"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SampleController struct {
	constants *bootstrap.Constants
}

func NewSampleController(constants *bootstrap.Constants) *SampleController {
	return &SampleController{
		constants: constants,
	}
}

func (sample *SampleController) Add(c *gin.Context) {
	type AddParams struct {
		Num1 int `uri:"num1" validate:"required,number,gt=10"`
		Num2 int `uri:"num2" validate:"required,number"`
	}
	param := controller.Validated[AddParams](c, &sample.constants.Context)

	sum := application_math.Add(param.Num1, param.Num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
