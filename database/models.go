package database

import (
	"time"
)

type User struct {
	ID             int    `json:"id"`
	LastName       string `json:"surname"`
	FirstName      string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	PassportSerie  string `json:"passport_serie"`
	PassportNumber string `json:"passport_number"`
}

type Task struct {
	ID        int       `json:"-"`
	TaskID    int       `json:"task_id"`
	UserID    int       `json:"-"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  string    `json:"duration"`
}
