package user

import (
	"gorm.io/gorm"
	t "todo/internal/task"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Tasks    []t.Task `gorm:"foreignKey:AuthorID,constraint:OnDelete:CASCADE"`
}
