package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
// Category 类别
type Category struct {
	Model
	Name  string `form:"name" json:"name" validate:"required" gorm:"column(name);"`
	Articles []Article `json:"articles" `
}

// GetALLCategory gets a list of articles based on paging constraints
func GetAllCategory() ([]*Category, error) {
	var categories []*Category

	err := db.Unscoped().Preload("Articles").Select("id,name").Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(categories)
	return categories, nil
}

// AddCategory 添加
func  AddCategory(c *Category) error {
	err:= db.Create(&c).Error
	return  err
}
// UpdateCategory 更新
func  UpdateCategory(id int , c *Category) error {
	err:= db.Unscoped().Model(&Category{}).Where("id = ? ", id).Update(c).Error
	return  err
}

// DelCategory 删除
func  DelCategoryById(id int) error {
	err:= db.Where("id = ?", id).Delete(&Category{}).Error
	return  err
}
