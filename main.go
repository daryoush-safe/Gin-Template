package main

import (
	"github.com/gin-gonic/gin"

	"first-project/src/bootstrap"
	"first-project/src/routes"
)

func main() {
	gin.DisableConsoleColor()
	ginEngine := gin.Default()

	var di = bootstrap.Run()

	routes.Run(ginEngine, di)

	ginEngine.Run(":8080")
}
