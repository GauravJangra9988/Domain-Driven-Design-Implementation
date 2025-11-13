package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github/gjangra9988/go-ddd/internal/user/application"
	"github/gjangra9988/go-ddd/internal/user/infrastructure/config"
	"github/gjangra9988/go-ddd/internal/user/infrastructure/persistence"
	"github/gjangra9988/go-ddd/internal/user/infrastructure/transport"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	log.Println("Connected to db")

	repo := persistence.NewUserRepo(db, redisClient)
	service := application.NewService(repo)
	handler := transport.NewHandler(service)

	router := gin.Default()
	handler.RegisterRoutes(router)

	router.Run(cfg.HTTPPort)
}