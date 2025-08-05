package auth

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID uint
	Token  string `gorm:"uniqueIndex"`
}
