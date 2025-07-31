package todo
import "time"

type Task struct {
    Id uint `gorm:"primaryKey"` 
    Title string `gorm:"not null"`
    Done bool
    CreatedAt time.Time
    UpdatedAt time.Time
}
