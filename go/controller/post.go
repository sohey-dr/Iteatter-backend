package controller

import (
	"iteatter/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	id := 1 // データベースでautoincrementする？
	title := c.PostForm("title")
	body := c.PostForm("body")
	// store data
	model.Post{
		Id:    id,
		title: title,
		Body:  body,
	}

}

func GetPosts(c *gin.Context) {
	//  select * from posts
	// posts := []string{"taro", "jiro", "ichiro"}
	posts := []model.Post{}
	user1 := model.Post{
		Id:       1,
		Title:    "this is title",
		Body:     "hey\nim john",
		UserFrom: 100,
	}
	posts = append(posts, user1)
	posts = append(posts, user1)
	posts = append(posts, user1)
	c.HTML(200, "list.html", gin.H{
		"posts": posts,
	})
}

func GetOnePost(c *gin.Context) {
	// select * from posts where id==??
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	}
	post := model.Post{
		Id:    id,
		Title: "this is title",
		Body:  "hey\nim john",
	}
	c.HTML(200, "show.html", gin.H{
		"post": post,
	})
}

// func UpdateOnePost(c *gin.Context) {
// 	var query model.Query
//
// }
