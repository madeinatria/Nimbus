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
		/*	python base - apis
			@app.get("/api/v1/user")
			@app.post("/api/v1/user/signup")
			@app.post("/api/v1/user/login")
			@app.post("/api/v1/client/signup")
			@app.get("/api/v1/client")
			@app.get("/api/v1/card/levels")
			@app.post("/api/v1/card/level/add")
			@app.post("/api/v1/card/add")
			@app.post("/api/v1/client/login")
			@app.post("/api/v1/authorize")
			@app.get("/api/v1/transaction")
			@app.get('/api/v1/cards')
		*/
	}
}
