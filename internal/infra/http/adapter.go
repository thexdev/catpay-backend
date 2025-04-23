package http

import (
	"catpay/internal/infra/http/handler"
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

	authHandler := handler.NewAuthHandler(h.db)

	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)

	return app
}
