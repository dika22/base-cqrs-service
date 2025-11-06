package service

import (
	"context"
	"cqrs-base/internal/domain"
	"cqrs-base/internal/query/repository"
	"cqrs-base/package/connection/cache"
)

type UserQueryService struct {
	repo repository.UserReadRepository
	cache cache.Cache
}

func NewUserQueryService(repo repository.UserReadRepository, cache cache.Cache) *UserQueryService {
	return &UserQueryService{repo : repo, cache : cache}
}

func (s *UserQueryService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserQueryService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAll(ctx)
}
