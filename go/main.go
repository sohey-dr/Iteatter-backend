package main

import (
	"iteatter/controller"
	"iteatter/infra"
	"iteatter/route"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("list", "templates/base.html", "templates/list.html")
	r.AddFromFiles("post", "templates/base.html", "templates/post.html")
	r.AddFromFiles("show", "templates/base.html", "templates/show.html")
	r.AddFromFiles("login", "templates/base.html", "templates/login.html")
	r.AddFromFiles("signup", "templates/base.html", "templates/signup.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createRender()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{})
	})
	router.GET("/post", func(c *gin.Context) {
		c.HTML(200, "post", gin.H{})
	})

	postEngine := router.Group("/posts")
	{
		postEngine.POST("/", controller.AddPost)
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

	infra.DbInit()
	router.Run(":8080")
}
