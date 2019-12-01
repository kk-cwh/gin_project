package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	Name     string    `json:"name" gorm:"column(name);"`
	Addr     string    `json:"addr" gorm:"column(addr);"`
	Age      int       `json:"age" gorm:"column(age);"`
	Birth    string    `json:"birth" gorm:"column(birth);"`
	Sex      int       `json:"sex" gorm:"column(sex);"`
	UpdateAt time.Time `json:"update_at" gorm:"column(update_at); description:"更新时间"`
	CreateAt time.Time `json:"create_at" gorm:"column(create_at);type(datetime)" description:"创建时间"`
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
