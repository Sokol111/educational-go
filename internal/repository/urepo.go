package repository

import (
	"context"
	"fmt"
	"github.com/Sokol111/educational-go/internal/model"
	"github.com/Sokol111/educational-go/internal/repository/db"
	"github.com/jackc/pgx/v5"
)

type UserPostgresRepository struct {
	queries db.Querier
}

func NewUserPostgresRepository(queries db.Querier) *UserPostgresRepository {
	return &UserPostgresRepository{queries}
}

func (r *UserPostgresRepository) GetById(ctx context.Context, id int64) (model.User, error) {
	if u, err := r.queries.GetUserById(ctx, id); err == nil {
		return toDomain(u), nil
	} else {
		return model.User{}, fmt.Errorf("failed to find user by id [%v], reason: %v", id, err)
	}
}

func (r *UserPostgresRepository) GetByName(ctx context.Context, name string) (found model.User, err error) {
	if u, err := r.queries.GetUserByName(ctx, name); err == nil {
		return toDomain(u), nil
	} else {
		return model.User{}, fmt.Errorf("failed to find user by name [%v], reason: %v", name, err)
	}
}

func (r *UserPostgresRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	if u, err := r.queries.CreateUser(ctx, db.CreateUserParams{Name: user.Name, Enabled: user.Enabled}); err == nil {
		return toDomain(u), nil
	} else {
		return model.User{}, fmt.Errorf("failed to create user [%v], reason: %v", user, err)
	}
}

func (r *UserPostgresRepository) Update(ctx context.Context, user model.User) (found model.User, err error) {
	if u, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{ID: user.ID, Name: user.Name, Version: user.Version, Enabled: user.Enabled}); err == nil {
		return toDomain(u), nil
	} else {
		if err == pgx.ErrNoRows {
			return model.User{}, fmt.Errorf("failed to update user [%v], reason: %v", user, "id not found or versions mismatch")
		}
		return model.User{}, fmt.Errorf("failed to update user [%v], reason: %v", user, err)
	}
}

func (r *UserPostgresRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	if users, err := r.queries.GetAllUsers(ctx); err == nil {
		s := make([]model.User, 0, len(users))
		for _, u := range users {
			s = append(s, toDomain(u))
		}
		return s, nil
	} else {
		return nil, fmt.Errorf("failed to get users, reason: %v", err)
	}
}

func toDomain(u db.User) model.User {
	return model.User{ID: u.ID, Version: u.Version, Enabled: u.Enabled, Name: u.Name, CreatedDate: u.CreatedDate.Time, LastModifiedDate: u.LastModifiedDate.Time}
}
