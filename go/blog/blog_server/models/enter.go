package models

import "time"

type Model struct {
	ID        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
