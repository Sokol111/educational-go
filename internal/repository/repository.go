package repository

import (
	"context"
	"github.com/Sokol111/educational-go/internal/model"
)

type UserRepository interface {
	GetById(ctx context.Context, id int64) (model.User, error)

	GetByName(ctx context.Context, name string) (model.User, error)

	Create(ctx context.Context, user model.User) (model.User, error)

	Update(ctx context.Context, user model.User) (model.User, error)

	GetUsers(ctx context.Context) ([]model.User, error)
}
