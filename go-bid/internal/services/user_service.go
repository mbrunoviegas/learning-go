package services

import (
	"context"
	"errors"
	"go-bid/internal/store/pgstore"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorDuplicatedEmailOrUsername = errors.New("username or email already exist")
	ErrorInvalidCredentials        = errors.New("invalid credentials")
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	input := pgstore.CreateUserParams{
		UserName:     userName,
		Email:        email,
		PasswordHash: hash,
		Bio:          bio,
	}

	id, err := us.queries.CreateUser(ctx, input)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, err
		}
	}

	return id, nil
}

func (us *UserService) AuthUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	user, err := us.queries.GetUserByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Error("user not found", "error", err)
			return uuid.UUID{}, ErrorInvalidCredentials
		}

		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		slog.Error("failed to match password", "error", err)

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, ErrorInvalidCredentials
		}

		return uuid.UUID{}, err
	}

	return user.ID, nil
}
