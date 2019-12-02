package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	Title  string `json:"title" gorm:"column(title);"`
	ContentMd  string `json:"content_md" gorm:"column(content_md);"`
	ContentHTML   string    `json:"content_html" gorm:"column(content_html);"`
	PageImageURL string `json:"page_image_url" gorm:"column(page_image_url);"`
	UserID   int    `json:"user_id" gorm:"column(user_id);"`
	CategoryID   int    `json:"category_id" gorm:"column(category_id);"`
	State   int    `json:"state" gorm:"column(state);"`
	PageView   int    `json:"page_view" gorm:"column(page_view);"`
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article

	err := db.Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(articles)
	return articles, nil
}
