package clinic_repository

import (
	"fmt"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

const (
	INSERT_CLINIC = `INSERT INTO CLINIC (name, cnpj, email_owner, cel_owner) VALUES ($1, $2, $3, $4)`
)

func (c ClinicRepository) Create(clinic *dto.ClinicDTORequest) *rest_err.RestErr {
	result, err := c.db.Exec(INSERT_CLINIC,
		clinic.Name,
		clinic.Cnpj,
		clinic.EmailOwner,
		clinic.CelOwner)

	if err != nil {
		return rest_err.NewInternalServeError("unable create clinic")
	}

	numberRowsEffected, err := result.RowsAffected()
	if err != nil {
		return rest_err.NewInternalServeError("unable create clinic")
	}

	fmt.Sprintf("foram afetadas %d linhas no clinic_repository/create.go", numberRowsEffected)

	return nil
}
