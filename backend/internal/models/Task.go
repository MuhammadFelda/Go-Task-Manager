package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectId		uint		`json:"project_id"`
	Issue			string		`json:"issue"`
	ProjectLeader	string		`json:"project_leader"`
	Communicator	[]int64		`gorm:"type:jsonb;serializer:json" json:"communicator"`
	Programmer		[]int64		`gorm:"type:jsonb;serializer:json" json:"programmer"`
	Designer		[]int64		`gorm:"type:jsonb;serializer:json" json:"designer"`
	Reviewer		[]int64		`gorm:"type:jsonb;serializer:json" json:"reviewer"`
	TicketLink		string		`json:"ticket_link"`
	RelatedLinks	[]string	`gorm:"type:jsonb;serializer:json" json:"related_links"`
	Description		string		`json:"description"`
	StartDate		*time.Time	`json:"start_date"`
	DueDate			*time.Time	`json:"due_date"`
	EndDate			*time.Time	`json:"end_date"`
	TimeUsed		float64		`json:"time_used"`
	IsActive		bool		`gorm:"default:true" json:"is_active"`
	IsAssign		bool		`gorm:"default:false" json:"is_assign"`
	Creator			uint		`json:"creator"`
	Updater			uint		`json:"updater"`

	Project		*Project	`gorm:"foreignKey:ProjectId" json:"project,omitempty"`
	Logtime		[]Logtime	`gorm:"foreignKey:TaskId" json:"logtime,omitempty"`
}