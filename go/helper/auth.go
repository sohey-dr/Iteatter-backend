package helper

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//ハッシュ化
func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

//ハッシュとマッチするかの確認
func PasswordValid(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

type Session struct {
	UserId interface{}
}

var LoginInfo Session
func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

			session := sessions.Default(c)
			LoginInfo.UserId = session.Get("UserId")

			// セッションがない場合、ログインフォームをだす
			if LoginInfo.UserId == nil {
					log.Println("ログインしていません")
					c.Redirect(http.StatusMovedPermanently, "/login")
					c.Abort() // これがないと続けて処理されてしまう
			} else {
					c.Set("UserId", LoginInfo.UserId) // ユーザidをセット
					c.Next()
			}
			log.Println("ログインチェック終わり")
	}
}