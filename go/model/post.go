package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `gorm:"size:255;not null"`
	Body  string `gorm:"size:2048;not null"`
}

type PostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
