package models

import (
	u "Projectmanagement_BE/utils"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Mail        string `json:"mail"`
	Bio         string `json:"bio"`
	UserID      uint   `json:"user_id"`
}

// Validate -
func (employee *Employee) Validate() (map[string]interface{}, bool) {

	if status, msg := u.CheckValidMail(employee.Mail); !status {
		return u.Message(status, msg), false
	}

	if status, msg := u.CheckValidPhone(employee.PhoneNumber); !status {
		return u.Message(status, msg), false
	}
	YourRootPassword	temp := &Employee{}

	//check for errors and duplicate emails
	err := GetDB().Table("employees").Where("mail = ?", employee.Mail).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Mail != "" {
		return u.Message(false, "Email address already in use by another employee."), false
	}

	// check for errors and duplicate phone nummbers
	err = GetDB().Table("employees").Where("phone_number = ?", employee.PhoneNumber).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.PhoneNumber != "" {
		return u.Message(false, "Phone number already in use by another employee."), false
	}

	return u.Message(false, "Employee requirement passed"), true
}
