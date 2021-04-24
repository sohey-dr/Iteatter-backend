package main

import (
	// postgres ドライバ

	"iteatter/controller"
	"iteatter/infra"
	"iteatter/route"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// TestUser : テーブルデータ
type TestUser struct {
	UserID   int
	Password string
}

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

// メイン関数
func main() {

	// // Db: データベースに接続するためのハンドラ
	// var Db *sql.DB
	// // Dbの初期化
	// Db, err := sql.Open("postgres", "host=postgres user=app_user password=password dbname=app_db sslmode=disable")
	// if err != nil {
	//     log.Fatal(err)
	// }
	//
	// // SQL文の構築
	// sql := "SELECT user_id, user_password FROM TEST_USER WHERE user_id=$1;"
	//
	// // preparedstatement の生成
	// pstatement, err := Db.Prepare(sql)
	// if err != nil {
	//     log.Fatal(err)
	// }
	//
	// // 検索パラメータ（ユーザID）
	// queryID := 1
	// // 検索結果格納用の TestUser
	// var testUser TestUser
	//
	// // queryID を埋め込み SQL の実行、検索結果1件の取得
	// err = pstatement.QueryRow(queryID).Scan(&testUser.UserID, &testUser.Password)
	// if err != nil {
	//     log.Fatal(err)
	// }
	//
	// // 検索結果の表示
	// fmt.Println(testUser.UserID, testUser.Password)

	router := gin.Default()
	router.HTMLRender = createRender()
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
