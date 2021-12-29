package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeteledAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Email     string `gorm:"not null;unique"`
	Password  string
}
