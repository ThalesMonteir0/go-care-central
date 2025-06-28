package postgres

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

type IUserRepository interface {
	FindUserByEmail(email string) (domain.User, error)
	FindUserByID(userID uuid.UUID) (domain.User, error)
	Update(user domain.User) error
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) FindUserByEmail(email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindUserByID(userID uuid.UUID) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(user domain.User) error {
	//TODO implement me
	panic("implement me")
}
