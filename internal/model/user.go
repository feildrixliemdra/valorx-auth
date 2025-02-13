package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" gorm:"primaryKey"`
	Name      string    `db:"name"`
	Email     string    `db:"email" gorm:"uniqueIndex:idx_email_deleted_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" gorm:"uniqueIndex:idx_email_deleted_at;default:null"`
}

func (u User) TableName() string {
	return "users"
}
