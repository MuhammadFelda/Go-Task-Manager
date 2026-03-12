package models

import "time"

type User struct {
	Id				uint		`json:"id"`
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	Role			string		`json:"role"`
	Password		string		`json:"password"`
	Avatar			string		`json:"avatar"`
	FaceEmbedding	float64		`json:"face_embedding,omitempty"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`

	Skills		[]Skill		`json:"skills,omitempty"`
	Logtimes	[]Logtime	`json:"logtimes,omitempty"`
	Logs		[]Log		`json:"logs,omitempty"`
}