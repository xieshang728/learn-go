package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlUrl = "%s:%s@tcp(%s:%s)/%s?charset=%s"

type User struct {
	Id      int64
	UserId  int64
	AddId   int64
	Name    string
	Address string
}

type Address struct {
	Id          int64
	UserId      int64
	AddId       int64
	AddName     string
	AddLocation string
}

func init() {

}

func main() {
	mysqlUrl := fmt.Sprintf(mysqlUrl, "root", "Aa123456", "39.96.5.21", "3306", "learn-go", "utf8")
	fmt.Println(mysqlUrl)
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if db.HasTable(&User{}) {
		db.AutoMigrate(&User{})
	} else {
		db.CreateTable(&User{})
	}

	if db.HasTable(&Address{}) {
		db.AutoMigrate(&Address{})
	} else {
		db.CreateTable(&Address{})
	}

	db.Model(&User{}).AddForeignKey("add_id", "addresses(id)", "RESTRICT", "RESTRICT")
	db.Model(&User{}).AddIndex("idx_user_add_id", "add_id")
	db.Model(&User{}).AddUniqueIndex("idx_user_id", "user_id")
}
