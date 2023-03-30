package apis

import "github.com/gin-gonic/gin"

func TempHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "home route"})
	return
}
