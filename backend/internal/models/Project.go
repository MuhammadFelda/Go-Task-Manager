package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name			string		`gorm:"not null" json:"name"`
	ProjectOwnerId	uint		`json:"project_owner_id"`
	Creator			uint		`json:"creator"`
	Updater			uint		`json:"updater"`
	IsDeleted		bool		`gorm:"not null" default:"false" json:"is_deleted"`

	ProjectOwner	*ProjectOwner	`gorm:"foreignKey:ProjectOwnerId" json:"project_owner,omitempty"`
	Task			[]Task			`gorm:"foreignKey:ProjectId" json:"tasks,omitempty"`
}