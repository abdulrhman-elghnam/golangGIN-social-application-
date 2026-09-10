package dto

type Login struct {
    Email    string `form:"email" json:"email" xml:"email" binding:"required,email"`
    Password string `form:"password" json:"password" xml:"password" binding:"required,min=8,max=100"`
}