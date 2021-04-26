package route

import (
	"fmt"
	"iteatter/model"
	"iteatter/helper"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//アカウント作成
func UserSignup(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordconf")
	if password != passwordConf {
		fmt.Println("パスワードが一致していません")
		ctx.Redirect(302, "/user/signup")
	}
	fmt.Println("username: " + username)
	fmt.Println("password: " + password)
	fmt.Println("passwordConf: " + passwordConf)

	//usernameが既に使われていないか確
	u := model.UserAlredy(username)
	if u != "" {
		fmt.Println("そのusernameは既に登録されています")
		ctx.Redirect(302, "/user/signup")
	}

	//passwordをハッシュ化
	hash, err := helper.PasswordHash(password)
	if err != nil {
		panic("hash err")
	}
	fmt.Println(hash)

	//登録
	model.UserInsert(username, hash)
	ctx.Redirect(302, "/")
}

//ログイン
func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println("username: " + username)
	fmt.Println("password: " + password)

	//username, passwordが正しいか確認
	err := model.UserCheck(username, password)
	if err == nil {
		fmt.Println("ログイン成功！！！！")
		// session := sessions.Default(ctx)
		// userId := model.SearchUserid(username)
    // session.Set("UserId", userId)
    // session.Save()
		// ctx.Redirect(302, "/")
	} else {
		fmt.Println("失敗！！！！！！")

		ctx.HTML(200, "login", gin.H{"message": "username または password が違います"})
	}

}
