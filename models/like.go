package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Like adalah representasi tabel likes di database
type Like struct {
	ID        string    `json:"id" gorm:"type:char(36);primaryKey"` // UUID untuk ID Like
	UserID    string    `json:"user_id"`                            // ID pengguna yang menyukai
	PostID    string    `json:"post_id"`                            // ID postingan yang disukai
	CreatedAt time.Time `json:"created_at"`                         // Waktu like dibuat
}

// BeforeCreate hook untuk menghasilkan UUID sebelum data like disimpan
func (l *Like) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New().String() // Generate UUID baru
	return
}
