package todo

import (
	"errors"
)

func AddTask(title string, isDone bool) error {
	if title == "" {
		return errors.New("title is required")
	}

	db, err := getDB()
	if err != nil {
		return err
	}

	task := Task{
		Title: title,
		Done:  isDone,
	}

	return db.Create(&task).Error
}

func ListTasks(showDone bool, showAll bool) ([]Task, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	q := db.Order("created_at DESC")
	switch {
	case showDone && showAll:
		return nil, errors.New("cannot show done and all tasks at the same time")
	case showDone:
		q = q.Where("done = ?", true)
	case showAll:
	default:
		q = q.Where("done = ?", false)
	}

	err = q.Find(&tasks).Error
	return tasks, err
}

func DoneTask(id string) error {
	db, err := getDB()

	if err != nil {
		return err
	}

	return db.Model(&Task{}).Where("id = ?", id).Update("done", true).Error
}

func DeleteTask(id string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	return db.Delete(&Task{}, id).Error
}

func ClearTasks() error {
	db, err := getDB()
	if err != nil {
		return err
	}

	return db.Where("1 = 1").Delete(&Task{}).Error
}

type TaskStats struct {
	Done      int
	Total     int
	PerDayAvg float64
}

func GetStats() (TaskStats, error) {
	db, err := getDB()
	if err != nil {
		return TaskStats{}, err
	}

	var stats TaskStats
	err = db.Raw(`
        SELECT
            COUNT(*) AS total,
            SUM(CASE WHEN done THEN 1 ELSE 0 END) AS done,
            (
                SELECT AVG(cnt) FROM (
                    SELECT DATE(created_at) as day, COUNT(*) as cnt
                    FROM tasks
                    WHERE done
                    GROUP BY day
                )
            ) AS per_day_avg
        FROM tasks
    `).Scan(&stats).Error

	return stats, err
}
