package auth

import (
	"errors"
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	SECRET_KEY = "34a5579cd861592331eff22e77c144832ec48492a05e13663ad02b42340ecb99"
)

func CreateJWT(user DTO.UserDTOResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"clinic_id": user.ClinicID,
		"name":      user.Name,
		"email":     user.Email,
		"exp":       time.Now().Add(8 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token invalido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token invalido")
	}

	clinicID := claims["clinic_id"].(string)

	return clinicID, err
}
