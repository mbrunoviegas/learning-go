package storage

import (
	"context"
	"log"
	"log/slog"
	"simple-crud/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/fx"
)

type UserDb struct {
	conn *pgx.Conn
}

func NewUserDb(lc fx.Lifecycle) *UserDb {
	conn, err := pgx.Connect(context.Background(), "postgres://root:root@localhost:5432/learning_go")

	if err != nil {
		slog.Error("unable to connect to database", "error", err)
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			slog.Info("Closing database connection...")
			return conn.Close(context.Background())
		},
	})

	return &UserDb{
		conn: conn,
	}
}

func (db *UserDb) AddUser(user models.User) (models.ID, error) {
	id := models.ID(uuid.New().String())
	_, err := db.conn.Exec(
		context.Background(),
		"INSERT INTO users (id, first_name, last_name, biography) VALUES ($1, $2, $3, $4)",
		id, user.FirstName, user.LastName, user.Biography,
	)

	return id, err
}

func (db *UserDb) GetUser(id models.ID) (models.User, bool) {
	var user models.User

	err := db.conn.QueryRow(
		context.Background(),
		"SELECT first_name, last_name, biography FROM users WHERE id = $1",
		id,
	).Scan(&user.FirstName, &user.LastName, &user.Biography)

	if err != nil {
		slog.Error("Failed to fetch user", "error", err)

		return user, false
	}

	return user, true
}

func (db *UserDb) ListUsers() []models.User {
	rows, err := db.conn.Query(context.Background(), "select first_name, last_name, biography FROM users")

	if err != nil {
		slog.Error("Failed to fetch users", "error", err)
		return nil
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Biography)
		if err != nil {
			slog.Error("Failed to scan user", "error", err)
			continue
		}
		users = append(users, user)
	}

	return users
}

func (db *UserDb) UpdateUser(id models.ID, updateData models.User) (models.User, bool) {
	_, err := db.conn.Exec(
		context.Background(),
		"UPDATE users SET first_name = $1, last_name = $2, biography = $3 WHERE id = $4",
		updateData.FirstName, updateData.LastName, updateData.Biography, id,
	)
	if err != nil {
		log.Printf("Failed to update user: %v\n", err)
		return updateData, false
	}

	return updateData, true
}

func (db *UserDb) DeleteUser(id models.ID) (models.User, bool) {
	var user models.User

	err := db.conn.QueryRow(
		context.Background(),
		"SELECT first_name, last_name, biography FROM users WHERE id = $1",
		id,
	).Scan(&user.FirstName, &user.LastName, &user.Biography)

	if err != nil {
		log.Printf("Failed to fetch user for deletion: %v\n", err)
		return user, false
	}

	_, err = db.conn.Exec(
		context.Background(),
		"DELETE FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		log.Printf("Failed to delete user: %v\n", err)
		return user, false
	}

	return user, true
}
