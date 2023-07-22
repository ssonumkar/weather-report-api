package auth

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ssonumkar/weather-report-api/internal/log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user User,logger log.CustomLogger) error {
	golangDateTime, _ := time.Parse("2006-01-02", user.DOB)
	stmt, err := r.db.Prepare("CALL RegisterUser(?, ?, ?)")
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to prepare statement: %s", err.Error()))
		return err
	}
	logger.Debug("Prepare successful")
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, golangDateTime)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to execute statement: %s", err.Error()))
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string, logger log.CustomLogger) (User, error) {
	var user User
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
