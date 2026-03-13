package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name			string		`gorm:"not null" json:"name"`
	Email			string		`gorm:"uniqueIndex;not null" json:"email"`
	Password		string		`gorm:"not null" json:"-"`	// Hide password from all json response
	Role			string		`gorm:"not null;default:'programmer'" json:"role"`
	Avatar			string		`json:"avatar"`
	FaceEmbedding	float64		`grom:"type:jsonb;serializer:json" json:"face_embedding,omitempty"`

	Skills			[]Skill		`gorm:"foreignKey:UserId" json:"skills,omitempty"`
	Logtimes		[]Logtime	`gorm:"foreignKey:UserId" json:"logtimes,omitempty"`
	Logs			[]Log		`gorm:"foreignKey:UserId" json:"logs,omitempty"`
}