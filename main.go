package main

import (
	"fmt"
	"gin_project/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "homestead:secret@127.0.0.1:33060/homestead?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println(err)
	defer db.Close()
	r := app.Init()
	r.Run(":8888")
}
