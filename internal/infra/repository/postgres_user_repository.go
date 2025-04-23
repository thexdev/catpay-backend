package repository

import (
	"catpay/internal/infra/repository/entity"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresUserRepository(conn *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{conn: conn}
}

// @todo replace with actual logic
func (r *PostgresUserRepository) Create(email, password, role string) error {
	ctx := context.Background()

	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	sql := `
		insert into users (
			uuid,
			email,
			password_hash,
			role
		) values ($1, $2, $3, $4)
	`

	uuid := uuid.NewString()

	if _, err := tx.Exec(ctx, sql, uuid, email, password, role); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *PostgresUserRepository) Exist(email string) error {
	sql := "select email from users where email = $1"

	var _email string
	err := r.conn.QueryRow(context.Background(), sql, email).Scan(&_email)
	if err == nil {
		return &entity.ErrUserAlreadyExist{}
	}

	return nil
}

func (r *PostgresUserRepository) GetHashedPasswordByEmail(
	email string,
) (string, error) {
	var passwordHash string

	sql := "select password_hash from users where email = $1"

	err := r.conn.QueryRow(context.Background(), sql, email).Scan(&passwordHash)
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}
