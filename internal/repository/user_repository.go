package repository

import (
	"clean-architecture-example/internal/domain"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

func (r *SQLiteUserRepository) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	if _, err := r.db.Exec(query); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	return nil
}

func (r *SQLiteUserRepository) Save(user *domain.User) error {
	query := `INSERT INTO users (id, name, email) VALUES (?, ?, ?)`
	if _, err := r.db.Exec(query, user.ID, user.Name, user.Email); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (r *SQLiteUserRepository) FindAll() ([]*domain.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating user rows: %w", err)
	}

	return users, nil
}
