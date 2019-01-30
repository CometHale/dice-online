package repository

import (
	"database/sql"
	"log"

	"github.com/comethale/dice-online/app/api/model/domain"

	"github.com/comethale/dice-online/app/api/shared/utils"
)

// UserRepository handles all direct db interaction for Users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates and returns a UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Get queries for and returns the User with the given id
func (repo *UserRepository) Get(id int64) (*domain.User, error) {

	stmt := `SELECT email, username FROM users WHERE id = $1`
	var email string
	var username string

	err := repo.db.QueryRow(stmt).Scan(&email, &username)

	if err != nil {
		log.Panicln(err)
	}

	return &domain.User{Email: email, Username: username, ID: id}, nil
}

// Update queries for the user with the given id and updates the row with the information given
func (repo *UserRepository) Update(email, password, username string, id int64) (*domain.User, error) {

	stmt := `UPDATE users SET email = $1, password = $2, username = $3 WHERE id = $4 RETURNING email, username`
	var hashedPassword string

	hashedPassword, err := utils.AuthHashPassword(password)

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	err = repo.db.QueryRow(stmt, email, hashedPassword, id).Scan(&email, &username)

	if err != nil {
		log.Panicln(err)
	}
	return &domain.User{Email: email, Username: username, ID: id}, nil
}

// Create creates a new User row in the db
func (repo *UserRepository) Create(email, password, username string) (*domain.User, error) {
	stmt := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	var hashedPassword string

	hashedPassword, err := utils.AuthHashPassword(password)

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	err = repo.db.QueryRow(stmt, email, username, hashedPassword).Scan(&id)

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return &domain.User{Email: email, ID: id, Username: username}, nil
}

// Delete removes and returns the User row with the given id
func (repo *UserRepository) Delete(id int64) (*domain.User, error) {

	stmt := `DELETE FROM users WHERE id = $1 RETURNING email`
	var email string

	err := repo.db.QueryRow(stmt).Scan(&email)

	if err != nil {
		log.Panicln(err)
	}

	return &domain.User{Email: email, ID: id}, nil
}
