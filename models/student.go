package models

import "time"

type Student struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Age       int       `gorm:"not null" json:"age"`
	Course    string    `gorm:"size:100" json:"course"`
	City      string    `gorm:"size:100" json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}