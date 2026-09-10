package dto

import "time"

type Signup struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required,alpha,min=2,max=50"`
	DOB      time.Time `form:"dob" json:"dob" xml:"dob" binding:"required"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=8,max=100"`
}