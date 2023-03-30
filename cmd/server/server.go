package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/madeinatria/Nimbus/cmd/apis"
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
			ctx.JSON(200, gin.H{"msg": "home route"})
		})

		// posts phonenumber and responds back with OTP
		api.POST("/initLogin", apis.TempHandler)

		// validates creds (phoneNumber and OTP) and responds back with auth token
		api.POST("/login", apis.TempHandler)

		// validates qr data i.e user info and for that resta
		// and triggers OTP for the user
		api.POST("/initiRedeemOffer", apis.TempHandler)

		// validates qr data and AMOUNT with OTP
		// Update db and send confirmation SMS
		api.POST("/offerRedeem", apis.TempHandler)

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
