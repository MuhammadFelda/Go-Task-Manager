package models

import "gorm.io/gorm"

type ProjectOwner struct {
	gorm.Model
	Name		string		`gorm:"not null" json:"name"`
	Creator		uint		`gorm:"not null" default:"false" json:"creator"`
	Updater		uint		`json:"updater"`
	IsDeleted	bool		`json:"is_deleted"`

	Project		[]Project	`gorm:"foreignKey:ProjectOwnerId" json:"projects,omitempty"`
}