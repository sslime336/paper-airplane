package orm

import (
	"time"

	"gorm.io/gorm"
)

func Models() []any {
	return []any{
		&Session{},
	}
}

type Model struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Session struct {
	Model

	OpenId          string `gorm:"uniqueIndex"`
	JsonizedHistory string
	TotalTokens     int64
}
