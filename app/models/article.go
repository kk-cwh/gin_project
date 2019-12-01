package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Name     string    `json:"name" gorm:"column(name);"`
	Addr     string    `json:"addr" gorm:"column(addr);"`
	Age      int       `json:"age" gorm:"column(age);"`
	Birth    string    `json:"birth" gorm:"column(birth);"`
	Sex      int       `json:"sex" gorm:"column(sex);"`
	UpdateAt time.Time `json:"update_at" gorm:"column(update_at); description:"更新时间"`
	CreateAt time.Time `json:"create_at" gorm:"column(create_at);type(datetime)" description:"创建时间"`
}

func (f *Article) TableName() string {
	return "articles"
}
