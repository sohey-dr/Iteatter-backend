package controller

import (
	"iteatter/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	//  select * from posts
	// posts := []string{"taro", "jiro", "ichiro"}
	posts := []model.Post{}
	posts = append(posts, model.Post{ // this is dummy
		Id:       1,
		Title:    "this is title",
		Body:     "hey\nim john",
		UserFrom: 100,
	})
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
