package models
import "time"

type Task struct {
    Id uint `gorm:"primaryKey"` 
    Title string `gorm:"not null"`
    Done bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

type TaskCreateInput struct {
    Title string `json:"title"`
    Done bool `json:"done"`
}

type TaskUpdateInput struct {
    Title string `json:"title"`
    Done bool `json:"done"`
}

type TaskStatsOutput struct {
	Done      int
	Total     int
	PerDayAvg float64
}
