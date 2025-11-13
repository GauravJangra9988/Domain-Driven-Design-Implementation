package application

import (
	"context"
	"github/gjangra9988/go-ddd/internal/user/domain/entities"
	Repositories "github/gjangra9988/go-ddd/internal/user/domain/repositories"
)

type UserRepository struct {
	repo Repositories.UserRepository
}

func NewService(r Repositories.UserRepository) *UserRepository {
	return &UserRepository{repo: r}
}

func (r *UserRepository) CreateUser(ctx context.Context, name, email string) (string, error){

	user := &entities.User{Name: name, Email: email}

	id, err := r.repo.Create(ctx,user)

	return id, err
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*UserResponse, error){

	user, err := r.repo.GetByID(ctx, id)

	userResponse := &UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	return userResponse, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, req UserUpdateRequest) (*UserResponse, error){

	user, err := r.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	err = r.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	userResponse := &UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return userResponse, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error{

	return r.repo.Delete(ctx, id)
}

func (r *UserRepository) CreateUserRedis(ctx context.Context, id, name, email string) (string, error){

	user := &entities.User{Name: name, Email: email}

	r.repo.RedisSetUser(ctx, id, user)

	return id, nil
}

func (r *UserRepository) GetUserRedis(ctx context.Context, id string) (*UserResponse, error){

	user, err := r.repo.RedisGetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	userResponse := &UserResponse{
		ID: id,
		Name: user.Name,
		Email: user.Email,
	}
	
	return userResponse, nil
}