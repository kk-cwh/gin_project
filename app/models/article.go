package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	Title  string `json:"title" validate:"required" gorm:"column(title);"`
	ContentMd  string `json:"content_md" validate:"required" gorm:"column(content_md);"`
	ContentHTML   string    `json:"content_html" validate:"required" gorm:"column(content_html);"`
	PageImageURL string `json:"page_image_url" validate:"required" gorm:"column(page_image_url);"`
	UserID   int    `json:"user_id" validate:"required" gorm:"column(user_id);"`
	User User ` json:"user" `
	CategoryID   uint    `json:"category_id" validate:"required" gorm:"column(category_id);"`
	Category Category ` json:"category" `
	State   int    `json:"state" validate:"required" gorm:"column(state);"`
	PageView   int    `json:"page_view"  gorm:"column(page_view);"`
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article

	err := db.Unscoped().Preload("Category").Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(articles)
	return articles, nil
}

func GetOneArticle( id int) (*Article, error) {
	var article Article
	//err := db.Model(&article).Where("id = ?",id).First(&article).Related(&article.User).Error
	err := db.Where("id=?", id).Preload("User").Preload("Category").First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(article)
	return &article, nil
}

func SaveArticle(m *Article) error {
	err := db.Create(&m).Error
	return err
}
