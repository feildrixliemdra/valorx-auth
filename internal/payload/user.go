package payload

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=5"`
	Email string `json:"email" binding:"required,email"`
}

type GetUserDetailData struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserRequest struct {
	ID    uuid.UUID `json:"-"`
	Name  string    `json:"name" binding:"required,min=5"`
	Email string    `json:"email" binding:"required,email"`
}
