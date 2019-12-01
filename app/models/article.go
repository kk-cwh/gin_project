package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	Name  string `json:"name" gorm:"column(name);"`
	Addr  string `json:"addr" gorm:"column(addr);"`
	Age   int    `json:"age" gorm:"column(age);"`
	Birth string `json:"birth" gorm:"column(birth);"`
	Sex   int    `json:"sex" gorm:"column(sex);"`
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article
	// db.LogMode(true)

	err := db.Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(articles)
	return articles, nil
}
