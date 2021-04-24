package infra

import (
	"fmt"
	"iteatter/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DbInit() *gorm.DB {
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		fmt.Errorf("cound not open database")
	}
	db.AutoMigrate(&domain.Post{})
	return db
}

func DbCreate(post domain.Post) {
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		fmt.Errorf("cound not open database")
	}
	db.Create(&post)
}

func DbRead(id ...int) []domain.Post {
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		fmt.Errorf("cound not open database")
	}
	var posts []domain.Post
	db.Find(&posts)
	return posts
}

func DbReadAll() []domain.Post {
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		fmt.Errorf("cound not open database")
	}
	var posts []domain.Post
	db.Find(&posts)
	return posts
}

func DbReadOne(id int) domain.Post {
	db, err := gorm.Open("sqlite3", "sample.db")
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	if err != nil {
		fmt.Errorf("cound not open database")
	}
	var post domain.Post
	db.First(&post, id)
	return post
}

func DbUpdate(id int, title string, body string) domain.Post {
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		fmt.Errorf("cound not open database")
	}
	var post domain.Post
	db.First(&post, id)
	post.Title = title
	post.Body = body
	db.Save(&post)
	return post
}

func DbDelete(id int) {
	// db, err := gorm.Open("postgres", "host=postgres user=app_user password=passwoerd dbname=app_db sslmode=disable")
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var post domain.Post
	db.First(&post, id)
	db.Delete(&post)
}
