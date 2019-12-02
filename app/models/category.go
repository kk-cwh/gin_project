package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
// Category 类别
type Category struct {
	Model
	Name  string `json:"name" binding:"required" gorm:"column(name);"`
}

// GetALLCategory gets a list of articles based on paging constraints
func GetAllCategory() ([]*Category, error) {
	var categories []*Category

	err := db.Unscoped().Select("id,name").Find(&categories).Error
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
// UpdateCategory UpdateCategory
func  UpdateCategory(id int , c *Category) error {
	fmt.Println(*c)
	err:= db.Unscoped().Model(&Category{}).Where("id = ? ", id).Update(c).Error
	return  err
}
