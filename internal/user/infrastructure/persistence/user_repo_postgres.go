package persistence

import (
	"context"
	"database/sql"
	"github/gjangra9988/go-ddd/internal/user/domain/entities"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, u *entities.User) (string, error){

	row := r.DB.QueryRowContext(ctx, `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, u.Name, u.Email)

	var id string

	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*entities.User, error){

	row := r.DB.QueryRowContext(ctx, `SELECT id, name, email FROM users WHERE id=$1`, id)

	user := &entities.User{}
	err := row.Scan(&user.ID,&user.Name,&user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, u *entities.User) error {

	_, err := r.DB.ExecContext(ctx, `UPDATE users SET name=$1, email=$2 WHERE id=$3`, u.Name, u.Email, u.ID)
	return err
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {

	_, err := r.DB.ExecContext(ctx, `DELETE FROM users WHERE id=$1`, id)
	return err
}