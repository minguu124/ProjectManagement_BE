package models

import "github.com/jinzhu/gorm"

// Task struct
type Task struct {
	gorm.Model
	Name        string `json:"name"`
	DateCreated string `json:"date_created"`
	DateEnd     string `json:"date_end"`
	Status      string `json:"task_status"`
	CreatorID   User   `json:"creator_id"`
	MemberID    []User `json:"member_ids"`
}
