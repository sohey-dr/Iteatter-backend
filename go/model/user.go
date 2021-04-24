package model

import (
	"iteatter/helper"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

//DBに接続
func dbOpen() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=mitei sslmode=disable")
	if err != nil {
		panic("データベース開けず！（dbOpen）")
	}
	return db
}

type User struct {
	gorm.Model
	Name     string
	Password string
}

//DB初期化
func UserInit() {
	db := dbOpen()
	db.AutoMigrate(&User{})
	defer db.Close()
}

//DB追加
func UserInsert(name string, hash string) {
	db := dbOpen()
	db.Create(&User{Name: name, Password: hash})
	defer db.Close()
}

//既にname	が登録されていないか確認
func UserAlredy(name string) string {
	db := dbOpen()
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Close()
	name := user.Name
	return name
}

//nameとpasswordが正しいか検証
func UserCheck(name string, password string) error {
	db := dbOpen()
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Close()

	err := helper.PasswordValid(user.Password, password)
	return err
}
