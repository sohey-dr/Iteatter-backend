package main

import (
	// postgres ドライバ

	"iteatter/controller"
	"iteatter/route"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// メイン関数
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})


	postEngine := router.Group("/posts")
	{
		// postEngine.POST("/", controller.AddPost)
		postEngine.GET("/", controller.GetPosts)
		postEngine.GET("/:id", controller.GetOnePost)
		// postEngine.PUT("/:id", controller.UpdateOnePost)
		// postEngine.DELETE("/:id", controller.DeleteOnePost)
	}


	router.GET("/login", route.Login)
	router.GET("/signup", route.Signup)
	user := router.Group("/user")
	{
		user.POST("/signup", route.UserSignup)
		user.POST("/login", route.UserLogin)
	}

	router.Run(":8080")
}
