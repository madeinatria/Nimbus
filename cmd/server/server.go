package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.New()
	setupApi(router)

	router.Run(fmt.Sprintf(":%s", "8081"))
}

func setupApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"supported_apis": "supported_endpoints"})
		})
	}
}
