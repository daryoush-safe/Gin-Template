package main

import (
	"github.com/gin-gonic/gin"

	"first-project/src/routes"
)

func main() {
	gin.DisableConsoleColor()
	ginEngine := routes.Run()

	ginEngine.Run(":8080")
}
