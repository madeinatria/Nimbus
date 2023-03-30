package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madeinatria/Nimbus/cmd/datastore"
	"github.com/madeinatria/Nimbus/cmd/model"
)

func TempHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})

}

func SignUp(ctx *gin.Context) {
	var payload model.SignUpPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if payload.Type == "client" {
		cli := datastore.Client{
			Name:    payload.Name,
			Phone:   payload.Phone,
			Details: payload.Details,
		}
		result := datastore.Db.Create(&cli)

		if result.Error != nil {
			// err
		}

		ctx.JSON(200, gin.H{"msg": "user created"})
		return
	}
	if payload.Type == "user" {
		cli := datastore.User{
			Name:    payload.Name,
			Phone:   payload.Phone,
			Details: payload.Details,
		}
		result := datastore.Db.Create(&cli)

		if result.Error != nil {
			// err
		}

		ctx.JSON(200, gin.H{"msg": "user created"})
		return

	}

	// ctx.JSON(200, gin.H{"msg": "home route"})
}

func InitLogin(ctx *gin.Context) {

	var payload model.InitLoginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// trigger otp
	// respond with uuid
	ctx.JSON(200, gin.H{"msg": "home route"})
}

func Login(ctx *gin.Context) {
	var payload model.LoginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// do validation and respond token
	ctx.JSON(200, gin.H{"msg": "home route"})
}

func InitRedeemOffer(ctx *gin.Context) {
	var payload model.InitOfferPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check if user exists and get user info
	// offer table - matching offers
	// if no offers - resp no offer and trigger sms for user
	// if yes then OTP
	ctx.JSON(200, gin.H{"msg": "home route"})
}

func RedeemOffer(ctx *gin.Context) {
	var payload model.InitOfferPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// find whcih offer - validate otp
	// add to transaction table - send confirmation
	ctx.JSON(200, gin.H{"msg": "home route"})
}
