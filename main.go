package main

import (
	"github.com/gin-gonic/gin"

	"first-project/routes"
)

func main() {
	gin.DisableConsoleColor()
	ginEngine := routes.Run()

	ginEngine.Run(":8080")
}
