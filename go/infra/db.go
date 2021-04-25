package infra

import (
	"iteatter/config"
	"iteatter/model"

	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	db.AutoMigrate(&model.Post{})
	return db
}

func DbCreate(post *model.Post) {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	db.Create(post)
}

func DbRead(id ...int) []model.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var posts []model.Post
	db.Find(&posts)
	return posts
}

func DbReadAll() []model.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var posts []model.Post
	db.Find(&posts)
	return posts
}

func DbReadOne(id int) model.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var post model.Post
	db.First(&post, id)
	return post
}

func DbUpdate(id int, title string, body string) model.Post {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("cound not open database")
	}
	var post model.Post
	db.First(&post, id)
	post.Title = title
	post.Body = body
	db.Save(&post)
	return post
}

func DbDelete(id int) {
	db, err := config.OpenDB()
	// db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("could not open database")
	}
	var post model.Post
	db.First(&post, id)
	db.Delete(&post)
}
