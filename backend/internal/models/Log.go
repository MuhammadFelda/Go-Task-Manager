package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	UserId			uint		`gorm:"not null" json:"user_id"`
	Target			string		`json:"target"`
	Description		string		`json:"description"`
}