package controller_v1

import (
	"first-project/application"
	"first-project/validation"

	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Add(c *gin.Context) {
	paramsValidation := validation.ValidateNunmber(c)
	if !paramsValidation {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters must be numeric"})
		return
	}

	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(c.Param("num2"))

	sum := application.Add(num1, num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
