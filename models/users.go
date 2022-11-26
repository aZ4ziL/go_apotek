package models

import "time"

// This is the configuration for the user model
type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	FirstName    string        `gorm:"size:20" json:"first_name"`
	LastName     string        `gorm:"size:20" json:"last_name"`
	Username     string        `gorm:"size:20;unique" json:"username"`
	Email        string        `gorm:"size:30;unique" json:"email"`
	Password     string        `gorm:"size:100" json:"-"`
	IsAdmin      bool          `gorm:"default:0" json:"-"`
	IsActive     bool          `gorm:"default:1" json:"is_active"`
	LastLogin    time.Time     `gorm:"null" json:"last_login"`
	DateJoined   time.Time     `gorm:"autoCreateTime" json:"date_joined"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
}
