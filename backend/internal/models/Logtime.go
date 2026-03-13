package models

import (
	"time"

	"gorm.io/gorm"
)

type Logtime struct {
	gorm.Model
	UserId		uint		`gorm:"not null" json:"user_id"`
	TaskId		uint		`gorm:"not null" json:"task_id"`
	Date		time.Time	`gorm:"not null" json:"date"`
	TimeUsed	float64		`gorm:"not null" json:"time_used"`
	Description	string		`json:"description"`

	User	*User	`gorm:"foreignKet:UserId" json:"user,omitempty"`
	Task	*Task	`gorm:"foreignKet:TaskId" json:"task,omitempty"`
}