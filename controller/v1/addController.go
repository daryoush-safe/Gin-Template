package controller_v1

import (
	"first-project/application"
	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"net/http"
)

type AddParams struct {
	Num1 string `uri:"num1" validate:"required,numeric"`
	Num2 string `uri:"num2" validate:"required,numeric"`
}

func Add(c *gin.Context) {
	var validate *validator.Validate = validator.New()
	// validate = validator.New()
	var params AddParams

	if err := c.ShouldBindUri(&params); err != nil {
		fmt.Print("ridi")
		return
	}
	if err := validate.Struct(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters must be numeric"})
		return
	}

	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(c.Param("num2"))

	sum := application.Add(num1, num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
