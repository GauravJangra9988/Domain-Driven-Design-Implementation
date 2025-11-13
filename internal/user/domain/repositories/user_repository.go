package Repositories

import (
	"context"
	"github/gjangra9988/go-ddd/internal/user/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, u *entities.User) (string, error)
	GetByID(ctx context.Context, id string) (*entities.User, error)
	Update(ctx context.Context, u *entities.User) error
	Delete(ctx context.Context, id string) error
	RedisSetUser(ctx context.Context, id string, u *entities.User) (string, error)
	RedisGetUser(ctx context.Context, id string) (*entities.User, error)
}