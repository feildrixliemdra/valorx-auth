package service

import (
	"context"
	"valorx-auth/internal/constant"
	"valorx-auth/internal/dto"
	"valorx-auth/internal/model"
	"valorx-auth/internal/payload"
	"valorx-auth/internal/repository"

	"github.com/google/uuid"
)

type IUserService interface {
	Create(ctx context.Context, p payload.CreateUserRequest) error
	GetByID(ctx context.Context, id uuid.UUID) (result payload.GetUserDetailData, err error)
	Update(ctx context.Context, request payload.UpdateUserRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type user struct {
	UserRepository repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &user{
		UserRepository: repo,
	}
}

func (s *user) Create(ctx context.Context, p payload.CreateUserRequest) (err error) {
	// 1. make sure no duplicate email
	existUser, err := s.UserRepository.GetBy(ctx, model.User{Email: p.Email})
	if err != nil {
		return err
	}

	if existUser != nil {
		return constant.ErrEmailAlreadyRegistered
	}

	// 2. transform create user request payload to user model
	usr := dto.CreateUserPayloadToUserModel(p)

	// 3. update user
	return s.UserRepository.Create(ctx, usr)
}

func (s *user) GetByID(ctx context.Context, id uuid.UUID) (result payload.GetUserDetailData, err error) {
	usr, err := s.UserRepository.GetBy(ctx, model.User{ID: id})
	if err != nil {
		return
	}

	if usr == nil {
		err = constant.ErrUserNotFound
		return
	}

	return dto.UserModelToUserDetailResponse(usr), nil
}

func (s *user) Update(ctx context.Context, p payload.UpdateUserRequest) error {

	// 1. make sure user exist
	currentUser, err := s.UserRepository.GetBy(ctx, model.User{ID: p.ID})
	if err != nil {
		return err
	}

	if currentUser == nil {
		return constant.ErrUserNotFound
	}

	// 2. transform request payload to user model
	usr := dto.UpdateUserPayloadToUserModel(p)

	// 3. update user
	return s.UserRepository.Update(ctx, usr)
}

func (s *user) Delete(ctx context.Context, id uuid.UUID) error {
	// 1. make sure user exist
	currentUser, err := s.UserRepository.GetBy(ctx, model.User{ID: id})
	if err != nil {
		return err
	}

	if currentUser == nil {
		return constant.ErrUserNotFound
	}

	// 2. delete user by id
	return s.UserRepository.DeleteByID(ctx, id)
}
