package controller_v1

import (
	"first-project/application"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Add(c *gin.Context) {
	num1, e := strconv.Atoi(c.Param("num1"))
	if e != nil {
		fmt.Println("invalid num1")
	}
	num2, e := strconv.Atoi(c.Param("num2"))
	if e != nil {
		fmt.Println("invalid num2")
	}
	sum := application.Add(num1, num2)
	c.String(http.StatusOK, strconv.Itoa(sum))
}
