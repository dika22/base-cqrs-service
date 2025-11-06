package repository

import (
	"context"
	"cqrs-base/internal/domain"

	"gorm.io/gorm"
)

type UserReadRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
}

type userReadRepository struct {
	db *gorm.DB
}

func NewUserReadRepository(db *gorm.DB) UserReadRepository {
	return &userReadRepository{db}
}

func (r *userReadRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user := domain.User{}
	if err := r.db.Table("users").Select("id", "name", "email", "created_at", "updated_at").
		Where("id = ?", id).First(&user).Error; err != nil {	
		return nil, err
	}
	return &user, nil
}

func (r *userReadRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Table("users").Select("id", "name", "email", "created_at", "updated_at").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
