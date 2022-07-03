package validator

import (
	"strconv"

	"golang_pet_project_1/internal/core/domain"
)

func ValidateUserAndPath(user *domain.User, id string) bool {
	return strconv.Itoa(user.ID) == id
}
