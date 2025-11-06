package repository

import (
	"context"
	"cqrs-base/internal/domain"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Exec(`INSERT INTO users (id, name, email, created_at, updated_at)
	 VALUES ($1, $2, $3, $4, $5)`,user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).Error; err != nil {
		return err
	 }
	return  nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Exec(`UPDATE users SET name=$1, email=$2, updated_at=$3 WHERE id=$4`,
		user.Name, user.Email, time.Now(), user.ID).Error; err != nil {
			return err
	}
	return  nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Exec(`DELETE FROM users WHERE id=$1`, id).Error; err != nil {
		return err
	}
	return nil
}
