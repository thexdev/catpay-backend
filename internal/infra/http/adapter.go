package http

import (
	"catpay/internal/app/service"
	"catpay/internal/infra/http/handler"
	"catpay/internal/infra/repository"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Http struct {
	db *pgxpool.Pool
}

func New() *Http {
	return &Http{}
}

func (h *Http) SetupDB() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:password@localhost:5431/catpay")
	if err != nil {
		log.Fatal("failed connect to database")
	}

	h.db = conn

	return conn
}

func (h *Http) Bootstrap() *fiber.App {
	h.SetupDB()

	app := fiber.New()

	userRepo := repository.NewPostgresUserRepository(h.db)
	passHasher := service.NewBcryptPasswordHasher()

	authHandler := handler.NewAuthHandler(userRepo, passHasher)

	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)

	return app
}
