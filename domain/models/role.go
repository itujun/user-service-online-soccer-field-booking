package models

import "time"

type Role struct {
	ID        uint   `gorm:"primaryKey;AutoIncrement"`
	Code      string `gorm:"varchar(15);not null"`
	Name      string `gorm:"varchar(20);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
