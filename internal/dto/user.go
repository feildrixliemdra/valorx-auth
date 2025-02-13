package dto

import (
	"valorx-auth/internal/model"
	"valorx-auth/internal/payload"

	"github.com/google/uuid"
)

func CreateUserPayloadToUserModel(p payload.CreateUserRequest) model.User {
	return model.User{
		ID:    uuid.New(),
		Name:  p.Name,
		Email: p.Email,
	}
}

func UpdateUserPayloadToUserModel(p payload.UpdateUserRequest) model.User {
	return model.User{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
	}
}

func UserModelToUserDetailResponse(u *model.User) payload.GetUserDetailData {
	return payload.GetUserDetailData{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
