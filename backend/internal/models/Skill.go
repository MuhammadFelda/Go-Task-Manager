package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	UserId		uint		`gorm:"not null" json:"user_id"`
	Name		string		`gorm:"not null" json:"skill"`

	User *User `gorm:"foreignKey:UserId" json:"users,omitempty"`
}