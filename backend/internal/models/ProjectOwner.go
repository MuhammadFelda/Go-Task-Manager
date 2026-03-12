package models

import "time"

type ProjectOwner struct {
	Id			uint		`json:"id"`
	Name		string		`json:"name"`
	Creator		uint		`json:"creator"`
	Updater		uint		`json:"updater"`
	IsDeleted	bool		`json:"is_deleted"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`

	Project		[]Project	`json:"project,omitempty"`
}