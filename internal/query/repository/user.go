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
	// row := r.db.QueryRowContext(ctx, `SELECT id, name, email, created_at, updated_at FROM users_read WHERE id = $1`, id)
	user := domain.User{}
	// if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
	// 	return nil, err
	// }
	return &user, nil
}

func (r *userReadRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	// rows, err := r.db.QueryContext(ctx, `SELECT id, name, email, created_at, updated_at FROM users_read ORDER BY created_at DESC`)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	var users []domain.User
	// for rows.Next() {
	// 	u := domain.User{}
	// 	if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, u)
	// }
	return users, nil
}
