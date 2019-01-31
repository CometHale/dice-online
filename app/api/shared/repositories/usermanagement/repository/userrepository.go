package repository

import (
	"database/sql"
	"errors"
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

	stmt := `SELECT email, username, highscore FROM users WHERE id = $1`
	var email string
	var username string
	var highscore int64

	err := repo.db.QueryRow(stmt, id).Scan(&email, &username, &highscore)

	if err != nil {
		log.Panicln(err)
	}

	return &domain.User{Email: email, Username: username, ID: id, HighScore: highscore}, nil
}

// GetPassword queries for and returns the password and id associated with the given Email
func (repo *UserRepository) GetPassword(email string) (string, int64, error) {
	// This kind of function should not exist in any production-ready application
	stmt := `SELECT password, id FROM users WHERE email = $1`
	var password string
	var id int64

	err := repo.db.QueryRow(stmt, email).Scan(&password, &id)

	if err != nil {
		log.Panicln(err)
	}

	return password, id, nil
}

// GetAll queries for and returns all Users
func (repo *UserRepository) GetAll() ([]*domain.User, error) {
	stmt := `SELECT email, username, highscore, id FROM users `
	var users []*domain.User

	rows, err := repo.db.Query(stmt)
	defer rows.Close()
	for rows.Next() {
		var email, username string
		var highscore, id int64
		rows.Scan(&email, &username, &highscore, &id)
		users = append(users, &domain.User{Email: email, Username: username, ID: id, HighScore: highscore})
	}

	if err != nil {
		log.Panicln(err)
	}

	return users, nil
}

// Update queries for the user with the given id and updates the row with the information given
func (repo *UserRepository) Update(email, username string, highscore, id int64) (*domain.User, error) {

	if id == 0 {
		log.Panicln("ID is a required value")
		return nil, errors.New("missing User ID for (*UserRepository).Update")
	}

	// only update if the update value isn't null, and prevent empty updates
	stmt := `UPDATE users SET email = COALESCE($1, email), ` +
		`username = COALESCE($2, username), ` +
		`highscore = COALESCE($3, highscore) ` +
		`WHERE id = $4 ` +
		`AND ($1 IS DISTINCT FROM email OR $2 IS DISTINCT FROM username OR $3 IS DISTINCT FROM highscore ) ` +
		`RETURNING email, username, highscore`

	err := repo.db.QueryRow(stmt, MakeDBString(email), MakeDBString(username), MakeDBInt(int(highscore)), id).Scan(&email, &username, &highscore)

	if err != nil {
		log.Panicln(err)
	}
	return &domain.User{Email: email, Username: username, ID: id, HighScore: highscore}, nil
}

// Create creates a new User row in the db
func (repo *UserRepository) Create(email, password, username string) (*domain.User, error) {
	stmt := `INSERT INTO users (email, username, highscore, password) VALUES ($1, $2, $3, $4) RETURNING id`

	var id, highscore int64

	hashedPassword, err := utils.AuthHashPassword(password)

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	err = repo.db.QueryRow(stmt, email, username, highscore, hashedPassword).Scan(&id)

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

// *****************************************************************************
//  Helper Functions
// *****************************************************************************

// MakeDBString converts a string into a db string
func MakeDBString(str string) sql.NullString {

	if str == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

// MakeDBInt converts an int into a db int
func MakeDBInt(val int) sql.NullInt64 {

	if val == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: int64(val),
		Valid: true,
	}
}
