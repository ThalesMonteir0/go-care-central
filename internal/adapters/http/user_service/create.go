package user_service

import (
	"go-care-central/internal/core/dto"
	"net/http"
)

func (u UserService) Create(response http.ResponseWriter, request *http.Request) {
	userDTO, err := dto.FromJsonUserDTORequest(request.Body)
	if err != nil {
		response.WriteHeader(err.Code)
		return
	}

	if err = u.useCase.Create(userDTO); err != nil {
		response.WriteHeader(err.Code)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
