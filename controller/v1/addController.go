package controller_v1

import (
	"first-project/application"
	"first-project/controller"

	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

type AddParams struct {
	Num1 string `uri:"num1" validate:"required,numeric,gt=10"`
	Num2 string `uri:"num2" validate:"required,numeric"`
}

func Add(c *gin.Context) {
	// paramsValidation, errorMsg := validation.ValidateNumber(c)
	// if !paramsValidation {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
	// 	return
	// }
	controller.Validated[AddParams](c)

	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(c.Param("num2"))

	sum := application.Add(num1, num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
