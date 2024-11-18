package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"title"`
	Email     string    `json:"content"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
