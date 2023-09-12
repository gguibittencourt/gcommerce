package database

import "time"

type Model struct {
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"->"`
	UpdatedAt time.Time `gorm:"->"`
}
