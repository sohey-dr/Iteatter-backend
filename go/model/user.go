package model

import (
  // "database/sql"
	// "fmt"
	// "log"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

//DBに接続
func DbOpen() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=mitei sslmode=disable")
	if err != nil {
		panic("データベース開けず！（dbOpen）")
	}
	return db
}