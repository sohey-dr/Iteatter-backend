package infra

import (
	"iteatter/config"
	"iteatter/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DbInit() *gorm.DB {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	db.AutoMigrate(&domain.Post{})
	return db
}

func DbCreate(post domain.Post) {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	db.Create(&post)
}

func DbRead(id ...int) []domain.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var posts []domain.Post
	db.Find(&posts)
	db.Close()
	return posts
}

func DbReadAll() []domain.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var posts []domain.Post
	db.Find(&posts)
	db.Close()
	return posts
}

func DbReadOne(id int) domain.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var post domain.Post
	db.First(&post, id)
	db.Close()
	return post
}

func DbUpdate(id int, title string, body string) domain.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var post domain.Post
	db.First(&post, id)
	post.Title = title
	post.Body = body
	db.Save(&post)
	db.Close()
	return post
}

func DbDelete(id int) {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("could not open database")
	}
	var post domain.Post
	db.First(&post, id)
	db.Delete(&post)
	defer db.Close()
}
