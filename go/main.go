package main

import (
	// postgres ドライバ
	"fmt"
	"iteatter/controller"
	"iteatter/model"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// メイン関数
func main() {
	rooter := gin.Default()
	rooter.LoadHTMLGlob("templates/*.html")
	rooter.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	postEngine := rooter.Group("/posts")
	{
		// postEngine.POST("/", controller.AddPost)
		postEngine.GET("/", controller.GetPosts)
		postEngine.GET("/:id", controller.GetOnePost)
		// postEngine.PUT("/:id", controller.UpdateOnePost)
		// postEngine.DELETE("/:id", controller.DeleteOnePost)
	}
	rooter.Run(":8080")
}
