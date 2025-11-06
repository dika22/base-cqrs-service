package service

import (
	"context"
	"cqrs-base/internal/domain"
	"cqrs-base/internal/query/repository"
)

type UserQueryService struct {
	repo repository.UserReadRepository
}

func NewUserQueryService(repo repository.UserReadRepository) *UserQueryService {
	return &UserQueryService{repo}
}

func (s *UserQueryService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserQueryService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAll(ctx)
}
