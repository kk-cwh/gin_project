package models

import "github.com/jinzhu/gorm"

type User struct {
	Model
	Username string `form:"username" json:"username" validate:"required" gorm:"column(username);"`
	Password string `form:"password" json:"password" validate:"required" gorm:"column(password);"`
	State    int    `json:"state" gorm:"column(state);"`
}

func SaveUser(u *User) error {
	err := db.Create(&u).Error
	return err
}

func FindByUserName(username string) (*User, error) {
	var user User
	err := db.Unscoped().Where("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return &user, err
}
