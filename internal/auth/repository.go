package auth

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ssonumkar/weather-report-api/internal/log"
)

// UserRepository handles database operations related to users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user User,logger log.CustomLogger) error {
	// golangDateTime := time.Now().Format(fmt.Sprintf("%d-%d-%d", user.DOB.Year, user.DOB.Month, user.DOB.Day)) 
	golangDateTime := time.Now().Format("2006-01-02 15:04:05") 
	stmt, err := r.db.Prepare("CALL RegisterUser(?, ?, ?)")
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to prepare statement: %s", err.Error()))
		return err
	}
	logger.Debug("Prepare successful")
	defer stmt.Close()
	// Execute the stored procedure with the user's data
	_, err = stmt.Exec(user.Username, user.Password, golangDateTime)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to execute statement: %s", err.Error()))
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string, logger log.CustomLogger) (User, error) {
	var user User
	// Perform the database query to retrieve the user by username
	query := "SELECT * FROM users WHERE username = ?"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.DOB)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(fmt.Sprintf("User not found: %s", err.Error()))
			return user, err
		}
		logger.Error(fmt.Sprintf("Failed to get user: %s", err.Error()))
		return user, err
	}
	return user, nil
}
