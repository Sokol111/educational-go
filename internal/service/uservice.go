package service

import (
	"context"
	"github.com/Sokol111/educational-go/internal/model"
	"github.com/Sokol111/educational-go/internal/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) GetById(ctx context.Context, id int64) (model.User, error) {
	return s.repository.GetById(ctx, id)
}

func (s *UserService) GetByName(ctx context.Context, name string) (model.User, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *UserService) Create(ctx context.Context, user model.User) (model.User, error) {
	return s.repository.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user model.User) (model.User, error) {
	return s.repository.Update(ctx, user)
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	return s.repository.GetUsers(ctx)
}
