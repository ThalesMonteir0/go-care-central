package clinic_service

import (
	"go-care-central/internal/core/dto"
	"net/http"
)

func (c ClinicService) Create(response http.ResponseWriter, request *http.Request) {
	body := request.Body
	clinicDTO, err := dto.FromJsonClinicRequest(body)
	if err != nil {
		response.WriteHeader(err.Code)
		return
	}

	err = c.useCase.Create(clinicDTO)
	if err != nil {
		response.WriteHeader(err.Code)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
