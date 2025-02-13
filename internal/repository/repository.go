package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
}

type Option struct {
	DB *gorm.DB
}

func InitiateRepository(opt Option) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(opt),
	}
}
