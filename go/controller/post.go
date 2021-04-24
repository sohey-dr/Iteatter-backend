package controller

import (
	"fmt"
	"iteatter/domain"
	"iteatter/infra"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	post := domain.Post{
		Title: c.PostForm("title"),
		Body:  c.PostForm("body"),
	}
	infra.DbCreate(&post)
	dst := fmt.Sprintf("/posts/%d", post.ID)
	fmt.Println(dst)
	fmt.Println(post.ID)
	c.Redirect(301, dst)
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
