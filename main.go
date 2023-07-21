package main

import (
	"dependencies/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.DefineRoutes(r)

	r.Run("localhost:8080")
}
