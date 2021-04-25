package route

import (
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.HTML(200, "login", gin.H{})
}

func Signup(ctx *gin.Context) {
	ctx.HTML(200, "signup", gin.H{})
}
