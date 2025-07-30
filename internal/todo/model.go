package todo
import "time"

type Task struct {
    Id int
    Title string
    Done bool
    CreatedAt time.Time
    UpdatedAt time.Time
}
