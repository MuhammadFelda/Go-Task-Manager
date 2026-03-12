package models

import "time"

type Logtime struct {
	Id			uint		`json:"id"`
	UserId		uint		`json:"user_id"`
	TaskId		uint		`json:"task_id"`
	Date		time.Time	`json:"date"`
	TimeUsed	float64		`json:"time_used"`
	Description	string		`json:"description"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`

	User	*User	`json:"user,omitempty"`
	Task	*Task	`json:"task,omitempty"`	
}