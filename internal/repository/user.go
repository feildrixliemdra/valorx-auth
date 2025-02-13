package repository

import (
	"context"
	"errors"
	"valorx-auth/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetBy(ctx context.Context, usr model.User) (*model.User, error)
	Update(ctx context.Context, user model.User) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type user struct {
	DB *gorm.DB
}

func NewUserRepository(opt Option) IUserRepository {
	return &user{
		DB: opt.DB,
	}
}

func (r *user) DeleteByID(ctx context.Context, id uuid.UUID) (err error) {
	result := r.DB.Model(&model.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"deleted_at": "NOW()",
		})

	return result.Error
}

func (r *user) Create(ctx context.Context, user model.User) (err error) {
	result := r.DB.WithContext(ctx).Create(&user)
	return result.Error
}

func (r *user) GetBy(ctx context.Context, usr model.User) (*model.User, error) {
	var result model.User

	query := r.DB.WithContext(ctx).
		Model(&model.User{}).
		Where("deleted_at IS NULL")

	if usr.Name != "" {
		query = query.Where("name ILIKE ?", "%"+usr.Name+"%")
	}

	if usr.Email != "" {
		query = query.Where("email = ?", usr.Email)
	}

	if usr.ID != uuid.Nil || usr.ID.String() != "00000000-0000-0000-0000-000000000000" {
		query = query.Where("id = ?", usr.ID)
	}

	err := query.First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *user) Update(ctx context.Context, user model.User) (err error) {
	result := r.DB.Model(&model.User{}).
		Where("id = ? AND deleted_at IS NULL", user.ID).
		Updates(map[string]interface{}{
			"email": user.Email,
			"name":  user.Name,
		})

	return result.Error
}
