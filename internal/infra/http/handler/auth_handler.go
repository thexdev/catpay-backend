package handler

import (
	"catpay/internal/app/service"
	"catpay/internal/app/usecase"
	"catpay/internal/infra/http/request"
	"catpay/internal/infra/repository"
	"catpay/internal/infra/repository/entity"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Validator = validator.New()
)

type AuthHandler struct {
	db *pgxpool.Pool

	loginUseCase    usecase.LoginUseCase
	registerUseCase usecase.RegisterUseCase
}

func NewAuthHandler(db *pgxpool.Pool) *AuthHandler {
	repo := repository.NewPostgresUserRepository(db)
	hasher := service.NewPasswordHasher()

	loginUseCase := usecase.NewLoginUseCase(repo, *hasher)
	registerUseCase := usecase.NewRegisterUseCase(repo, *hasher)

	return &AuthHandler{
		db: db,

		// all use cases...
		loginUseCase:    *loginUseCase,
		registerUseCase: *registerUseCase,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := request.NewLoginRequest()

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if ok, err := req.Validate(); !ok {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": req.Errors(err),
		})
	}

	usecase := h.loginUseCase.SetCredential(
		usecase.LoginRequest{
			Email:    req.Email,
			Password: req.Password,
		},
	)

	ok, err := usecase.Execute()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "account not found",
		})
	}

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid user credential",
		})
	}

	return c.JSON(req)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	req := request.NewRegisterRequest()

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if ok, err := req.Validate(); !ok {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": req.Errors(err),
		})
	}

	usecase := h.registerUseCase.SetCredential(usecase.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	_, err := usecase.Execute()
	fmt.Println(err)

	if err != nil {
		if errors.Is(err, &entity.ErrUserAlreadyExist{}) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "user already exist",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{})
}
