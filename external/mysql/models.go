package mysql

import (
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
	Login_as string `gorm:"size:255;not null"`
}
