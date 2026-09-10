package dto

type Signup struct {
    User     string `form:"user" json:"user" xml:"user" binding:"required,min=3,max=20"`
    Name     string `form:"name" json:"name" xml:"name" binding:"required,alpha,min=2,max=50"`
    Age      int    `form:"age" json:"age" xml:"age" binding:"required,min=18,max=100"`
    Password string `form:"password" json:"password" xml:"password" binding:"required,min=8,max=100"`
}
