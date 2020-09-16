package models

import (
	u "Projectmanagement_BE/utils"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User struct
type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Employee Employee  `json:"employee"`
	Projects []Project `gorm:"many2many:user_projects"`
	Roles    []Role    `gorm:"`
	Token    string    `json:"token";sql:"-"`
}

// Create -
func (user *User) Create() map[string]interface{} {
	if msg, status := user.Validate(); !status {
		return msg
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)
	GetDB().Create(user.Employee)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered user
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

// Login -
func Login(name, password string) map[string]interface{} {

	user := &User{}
	if user, status := GetUserByName(name); !status {
		return u.Message(false, "Connection error. Please retry.")
	}
	if user == nil {
		return u.Message(false, "User name not found.")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Wrong password.")
	}

	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

// GetUserByID -
func GetUserByID(id uint) (*User, bool) {
	user := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}

	user.Password = ""
	return user, true
}

// GetUserByName -
func GetUserByName(name string) (*User, bool) {
	user := &User{}
	err := GetDB().Table("users").Where("name = ?", name).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return user, true
}

// Validate - check user info
func (user *User) Validate() (map[string]interface{}, bool) {

	if msg, status := user.Employee.Validate(); !status {
		return msg, status
	}

	temp := &User{}

	//check for errors and duplicate user name
	err := GetDB().Table("users").Where("name = ?", user.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry."), false
	}
	if temp.Name != "" {
		return u.Message(false, "User name already in use by another user."), false
	}

	return u.Message(false, "Requirement passed."), true
}
