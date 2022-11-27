package models

import (
	"time"

	"github.com/aZ4zil/go_apotek/auth"
)

// This is the configuration for the user model
type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	FirstName    string        `gorm:"size:20" json:"first_name"`
	LastName     string        `gorm:"size:20" json:"last_name"`
	Username     string        `gorm:"size:20;unique" json:"username"`
	Email        string        `gorm:"size:30;unique" json:"email"`
	Password     string        `gorm:"size:150" json:"-"`
	IsAdmin      bool          `gorm:"default:0" json:"-"`
	IsActive     bool          `gorm:"default:1" json:"is_active"`
	LastLogin    time.Time     `gorm:"null" json:"last_login"`
	DateJoined   time.Time     `gorm:"autoCreateTime" json:"date_joined"`
	Transactions []Transaction `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transactions"`
}

// CreateNewUser
// create new user
func CreateNewUser(user *User) error {
	db = Connection()
	user.Password = auth.EncryptionPassword(user.Password)
	return db.Create(user).Error
}

// GetUserByEmail
// return user by given param `email`
func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("email = ?", email).Preload("Transactions").First(&user).Error
	return user, err
}
