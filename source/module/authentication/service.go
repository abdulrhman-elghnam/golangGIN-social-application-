package authentication

import (
	"errors"
	"golang/source/common/jwt"
	"golang/source/common/security"
	"golang/source/database"
	"golang/source/database/model"
	"golang/source/module/authentication/dto"
	"os"

	"github.com/gin-gonic/gin"
)

type ServiceError struct {
	Status  int
	Message string
}

func SignupHandler(ctx *gin.Context) (map[string]any, *ServiceError) {
	data := ctx.MustGet("body").(*dto.Signup)

	users, err := database.UserRepository.Find("email", data.Email)

	if err != nil {
		return nil, &ServiceError{Status: 500, Message: err.Error()}
	}

	if len(users) > 0 {
		return nil, &ServiceError{Status: 409, Message: "email already exists"}
	}

	hashedPassword, err := security.HashPassword(data.Password)
	if err != nil {
		return nil, &ServiceError{Status: 500, Message: "could not secure password"}
	}

	user := model.User{
		Name:     data.Name,
		DOB:      data.DOB,
		Email:    data.Email,
		Password: hashedPassword,
	}

	err = database.UserRepository.Create(user)

	if err != nil {
		return nil, &ServiceError{Status: 500, Message: err.Error()}
	}

	return map[string]any{"email": user.Email}, nil
}

func LoginHandler(ctx *gin.Context) (map[string]any, *ServiceError) {
	data := ctx.MustGet("body").(*dto.Login)
	users, err := database.UserRepository.Find("email", data.Email)
	if err != nil {
		return nil, &ServiceError{Status: 500, Message: err.Error()}
	}
	if len(users) == 0 || !security.ComparePassword(data.Password, users[0].Password) {
		return nil, &ServiceError{Status: 401, Message: "invalid email or password"}
	}

	secret := os.Getenv("TKN_KEY")
	if secret == "" {
		return nil, &ServiceError{Status: 500, Message: "token secret is not configured"}
	}
	token, err := jwt.GenerateToken(users[0].ID, secret)
	if err != nil {
		return nil, &ServiceError{Status: 500, Message: "could not create token"}
	}

	return map[string]any{"token": token, "user": map[string]any{
		"id": users[0].ID, "name": users[0].Name, "email": users[0].Email,
	}}, nil
}

var _ = errors.New
