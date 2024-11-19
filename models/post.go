package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Post adalah representasi tabel di database
type Post struct {
	ID        string    `json:"id" gorm:"type:char(36);primaryKey"` // Gunakan UUID sebagai ID
	Title     string    `json:"title"`                              // Judul postingan
	Content   string    `json:"content"`                            // Konten postingan
	CreatedAt time.Time `json:"created_at"`                         // Waktu saat postingan dibuat
	UpdatedAt time.Time `json:"updated_at"`                         // Waktu saat postingan diperbarui
}

// BeforeCreate hook untuk menghasilkan UUID sebelum data disimpan
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String() // Generate UUID baru
	return
}
