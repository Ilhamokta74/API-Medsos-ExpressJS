package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User adalah representasi tabel users di database
type User struct {
	ID        string    `json:"id" gorm:"type:char(36);primaryKey"` // UUID untuk ID User
	Name      string    `json:"name"`                               // Nama user
	Email     string    `json:"email" gorm:"unique"`                // Email user (unik)
	Password  string    `json:"-"`                                  // Password user (disembunyikan di respon)
	CreatedAt time.Time `json:"created_at"`                         // Waktu user dibuat
	UpdatedAt time.Time `json:"updated_at"`                         // Waktu user diperbarui
}

// BeforeCreate hook untuk menghasilkan UUID sebelum data user disimpan
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String() // Generate UUID baru
	return
}
