package models

import "time"

type Skill struct {
	Id			uint		`json:"id"`
	UserId		uint		`json:"user_id"`
	Name		string		`json:"skill"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`

	User *User `json:"user,omitempty"`
}