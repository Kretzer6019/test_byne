package router

import (
	"dependencies/controllers/clients"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r *gin.Engine) {
	r.GET("/pares", clients.GetEven)
	r.GET("/impares", clients.GetOdd)
}
