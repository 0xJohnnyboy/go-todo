package task 

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
    Title string `gorm:"not null"`
    Done bool
	AuthorId uint
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
