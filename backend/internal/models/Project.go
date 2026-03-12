package models

import "time"

type Project struct {
	Id				uint		`json:"id"`
	Name			string		`json:"name"`
	ProjectOwnerId	uint		`json:"project_owner_id"`
	Creator			uint		`json:"creator"`
	Updater			uint		`json:"updater"`
	IsDeleted		bool		`json:"is_deleted"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`

	ProjectOwner	*ProjectOwner	`json:"project_owner,omitempty"`
	Task			[]Task			`json:"task,omitempty"`
}