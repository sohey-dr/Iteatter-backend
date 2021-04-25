package controller

import (
	"iteatter/infra"
	"iteatter/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiAddPost(c *gin.Context) {
	var input model.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := model.Post{
		Title: input.Title,
		Body:  input.Body,
	}
	infra.DbCreate(&post)
	c.JSON(http.StatusOK, post)
}
func ApiGetPosts(c *gin.Context) {
	posts := infra.DbReadAll()
	c.JSON(http.StatusOK, posts)
}

func ApiGetOnePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	post := infra.DbReadOne(id)
	c.JSON(http.StatusOK, post)
}
