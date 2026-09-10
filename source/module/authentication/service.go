package authentication

import (
	"golang/source/common/structure"
	"golang/source/database"
	"golang/source/database/model"
	"golang/source/module/authentication/dto"

	"github.com/gin-gonic/gin"
)

func SignupHandler(ctx *gin.Context) structure.Response {
	data := ctx.MustGet("body").(*dto.Signup)

user := model.User{
	Name:     data.Name,
	DOB:      data.DOB,
	Email:    data.Email,
	Password: data.Password,
}


	err := database.UserRepository.Create(user)

	if err != nil {
		return structure.Response{
			Success: false,
			Error: &structure.ErrorInfo{
				Message: err.Error(),
			},
		}
	}

	return structure.Response{
		Success: true,
		Result:  "User created successfully",
	}
}


func LoginHandler(ctx *gin.Context) structure.Response {
	data := ctx.MustGet("body").(*dto.Signup)

	user := model.User{
		Email:    data.Email,
		Password: data.Password,
	}
	// err := database.UserRepository.FindByID()

	err := database.UserRepository.Create(user)

	if err != nil {
		return structure.Response{
			Success: false,
			Error: &structure.ErrorInfo{
				Message: err.Error(),
			},
		}
	}

	return structure.Response{
		Success: true,
		Result:  "User created successfully",
	}
}
