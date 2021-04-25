package controller

import (
	"fmt"
	"iteatter/infra"
	"iteatter/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	post := model.Post{
		Title: c.PostForm("title"),
		Body:  c.PostForm("body"),
	}
	infra.DbCreate(&post)
	dst := fmt.Sprintf("/posts/%d", post.ID)
	c.Redirect(301, dst)
}

func GetPosts(c *gin.Context) {
	posts := infra.DbReadAll()
	c.HTML(200, "list", gin.H{
		"posts": posts,
	})
}

type error struct {
	Text string
}

func GetOnePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	}
	post := infra.DbReadOne(id)
	c.HTML(200, "show", gin.H{
		"post": post,
	})
}

// func UpdateOnePost(c *gin.Context) {
// 	var query model.Query
//
// }
