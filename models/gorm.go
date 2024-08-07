package models

import (
	"time"

	"gorm.io/gorm"
)

type GORMModel interface {
}

type MavisModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
