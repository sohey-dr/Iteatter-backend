package route

import (
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{})
}

func Signup(ctx *gin.Context) {
	ctx.HTML(200, "signup.html", gin.H{})
}
