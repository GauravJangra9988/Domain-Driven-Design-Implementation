package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github/gjangra9988/go-ddd/internal/user/domain/entities"

	"github.com/go-redis/redis/v8"
)

type UserRepo struct {
	DB *sql.DB
	RedisDB *redis.Client
}

func NewUserRepo(db *sql.DB, redisClient *redis.Client) *UserRepo {
	return &UserRepo{
		DB: db,
		RedisDB: redisClient,
	}
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

func (r *UserRepo) RedisSetUser(ctx context.Context, id string, u *entities.User) (string, error){

	val := map[string]string{
		"name": u.Name,
		"email": u.Email,
	}

	redisVal, err := json.Marshal(&val)
	if err != nil {
		return "", err
	}

	err = r.RedisDB.Set(ctx, id, redisVal, 0).Err()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepo) RedisGetUser(ctx context.Context, id string) (*entities.User, error){
	
	user := &entities.User{}

	key, err := r.RedisDB.Get(ctx, id).Result()
	if err == redis.Nil {
		return nil, errors.New("key not found")
	}

	err = json.Unmarshal([]byte(key), &user)
	if err != nil {
		return nil, errors.New("invalid data in cache")
	}

	return user, nil
}