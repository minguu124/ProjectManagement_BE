package models

import (
	u "Projectmanagement_BE/utils"
	"time"

	"github.com/jinzhu/gorm"
)

// Permission struct -
type Permission struct {
	gorm.Model
	List []uint
}

// Role struct -
type Role struct {
	gorm.Model
	Permissions []Permission
	ProjectID   uint
}

// Project struct
type Project struct {
	gorm.Model
	Title       string
	DateStarted string
	Tasks       []Task `gorm:"foreignKey:project_id;association_foreignkey:id"`
	Logs        []Log  `gorm:"foreignKey:project_id;association_foreignkey:id"` // log user activities
	Roles       []Role `gorm:"foreignKey:project_id;association_foreignkey:id`  // user who creates project
	CreatorID   uint
}

// ProjectUser struct - project user relation
type ProjectUser struct {
	UserID    uint
	ProjectID uint
	CreatedAt time.Time
}

// Create project
func (project *Project) Create() map[string]interface{} {

	GetDB().Create(project)

	if project.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered user

	response := u.Message(true, "Project has been created")
	response["project"] = project
	return response
}
