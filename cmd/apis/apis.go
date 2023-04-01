package apis

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madeinatria/Nimbus/cmd/datastore"
	"github.com/madeinatria/Nimbus/cmd/model"
	nimbusUtils "github.com/madeinatria/Nimbus/cmd/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": result.Error.Error()})
			return
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
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"msg": "user created"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid request"})
}

func InitLogin(ctx *gin.Context) {

	var payload model.InitLoginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var client datastore.Client
	findPhErr := datastore.Db.Model(&datastore.Client{}).Where("phone = ?", payload.Phone).First(&client)
	if errors.Is(findPhErr.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "phone number not found"})
		return
	}

	log.Println("identified user : ", client.Name)

	// genOtp, errOtp := nimbusUtils.GenerateOTP()
	// if errOtp != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in otp gen"})
	// 	return
	// }
	genOtp := "1100"

	genNonce, errNonce := nimbusUtils.GenerateNonce()
	if errNonce != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in nonce gen"})
		return
	}

	// persist OTP
	// result := datastore.Db.Clauses(db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
	//   }).Create(&users)).Create(&datastore.KeyPair{
	// 	Nonce: genNonce,
	// 	Key:   client.Phone,
	// 	Value: genOtp,
	// })

	result := datastore.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "nonce"}),
	}).Create(&datastore.KeyPair{
		Nonce: genNonce,
		Key:   client.Phone,
		Value: genOtp,
	})

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not initLogin"})
		return
	}
	ctx.JSON(200, gin.H{"nonce": genNonce})
}

func Login(ctx *gin.Context) {
	var payload model.LoginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var OTPEntry datastore.KeyPair
	result := datastore.Db.Model(&datastore.KeyPair{}).Where("nonce = ?", payload.Nonce).First(&OTPEntry)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "otp entry not found"})
		return
	}

	if OTPEntry.Key == payload.Phone && OTPEntry.Value == payload.OTP {
		ctx.JSON(http.StatusOK, gin.H{"msg": "logged in"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid state"})
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

func OffersHandler(ctx *gin.Context) {

	if ctx.Request.Method == http.MethodGet {
		var matchingOffers []datastore.Offers
		// result := datastore.Db.Model(&datastore.Offers{}).Find(&matchingOffers)
		// result := datastore.Db.Table("offers").Where("").Find(&matchingOffers)
		// if result.Error != nil {

		// }
		ctx.JSON(http.StatusOK, gin.H{"data": matchingOffers})
		return
	}
	if ctx.Request.Method == http.MethodPost {

		var off []datastore.Offers
		datastore.Db.Table("offers").Where("").Find(&off)
		// 	datastore.Offers{
		// 		CardName:     datastore.Card(datastore.Basic),
		// 		DiscountRate: 10,
		// 		ClientId:     0,
		// 	},
		// )
		// log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{"data": "created"})
		return
	}
	if ctx.Request.Method == http.MethodPatch {

	}
	if ctx.Request.Method == http.MethodDelete {

	}

}
