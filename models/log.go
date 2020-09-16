package models

import (
	"github.com/jinzhu/gorm"
)

// Log struct
type Log struct {
	gorm.Model
	Activity string `json:"activity"`
	Task     Task   `json:"task"`
	User     User   `json:"user"`
}
