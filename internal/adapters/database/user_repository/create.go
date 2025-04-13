package user_repository

import (
	"fmt"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

const (
	INSERT_USER = `INSERT INTO USER (name, email, password, cel, clinic_id) VALUES ($1, $2, $3, $4, $5)`
)

func (u UserRepository) Create(user *dto.UserDTORequest) *rest_err.RestErr {
	result, err := u.db.Exec(INSERT_USER,
		user.Name,
		user.Email,
		user.Password,
		user.ClinicID)
	if err != nil {
		return rest_err.NewBadRequestError("unable create user")
	}

	numberRowsEffected, err := result.RowsAffected()
	if err != nil {
		return rest_err.NewBadRequestError("unable create user")
	}

	fmt.Sprintf("numero de linhas afetadas: %d em user_repository/create.go", numberRowsEffected)

	return nil
}
