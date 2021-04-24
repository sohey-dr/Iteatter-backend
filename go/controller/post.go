package controller

import (
	"iteatter/domain"
	"iteatter/infra"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	infra.DbCreate(domain.Post{
		Title: c.PostForm("title"),
		Body:  c.PostForm("body"),
	})
	c.Redirect(301, "/")
}

func GetPosts(c *gin.Context) {
	posts := infra.DbReadAll()
	c.HTML(200, "list", gin.H{
		"posts": posts,
	})
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
