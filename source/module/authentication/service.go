package authentication

import "golang/source/module/authentication/dto"

func Signup(data dto.Signup) (string, error) {
	return "User created successfully", nil
}