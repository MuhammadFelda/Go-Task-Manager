package models

import "time"

type Task struct {
	Id				uint		`json:"id"`
	ProjectId		uint		`json:"project_id"`
	Issue			string		`json:"issue"`
	ProjectLeader	string		`json:"project_leader"`
	Communicator	[]int64		`json:"communicator"`
	Programmer		[]int64		`json:"programmer"`
	Designer		[]int64		`json:"designer"`
	Reviewer		[]int64		`json:"reviewer"`
	TicketLink		string		`json:"ticket_link"`
	RelatedLinks	[]string	`json:"related_links"`
	Description		string		`json:"description"`
	StartDate		*time.Time	`json:"start_date"`
	DueDate			*time.Time	`json:"due_date"`
	EndDate			*time.Time	`json:"end_date"`
	TimeUsed		float64		`json:"time_used"`
	IsActive		bool		`json:"is_active"`
	IsAssign		bool		`json:"is_assign"`
	Creator			uint		`json:"creator"`
	Updater			uint		`json:"updater"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`

	Project		*Project	`json:"project,omitempty"`
	Logtime		[]Logtime	`json:"logtime,omitempty"`
}