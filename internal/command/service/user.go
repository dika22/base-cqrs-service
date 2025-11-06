package service

import (
	"context"
	"cqrs-base/internal/command/repository"
	"cqrs-base/internal/domain"
	"fmt"
	"time"
)

type UserService struct {
	repo      repository.UserRepository
	// publisher EventPublisher
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, id, name, email string) error {
	user := &domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return err
	}

	event := domain.UserEvent{
		EventType: "UserCreated",
		UserID:    id,
		Name:      name,
		Email:     email,
		Time:      time.Now(),
	}
	// return s.publisher.Publish(ctx, event)
	fmt.Println("debug", event)
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, id, name, email string) error {
	user := &domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		UpdatedAt: time.Now(),
	}
	if err := s.repo.Update(ctx, user); err != nil {
		return err
	}

	event := domain.UserEvent{
		EventType: "UserUpdated",
		UserID:    id,
		Name:      name,
		Email:     email,
		Time:      time.Now(),
	}
	// return s.publisher.Publish(ctx, event)

	fmt.Println("debug", event)
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	event := domain.UserEvent{
		EventType: "UserDeleted",
		UserID:    id,
		Time:      time.Now(),
	}
	// return s.publisher.Publish(ctx, event)

	fmt.Println("debug", event)

	return nil
}
