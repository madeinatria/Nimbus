package apis

import "github.com/gin-gonic/gin"

func TempHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})

}

func InitLogin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})
}

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})
}

func InitRedeemOffer(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})
}

func RedeemOffer(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})
}
