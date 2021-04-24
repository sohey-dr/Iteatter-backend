package domain

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body  string `gorm:"size:2048"`
}
