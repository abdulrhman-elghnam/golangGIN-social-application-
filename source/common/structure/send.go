package structure

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func OK(ctx *gin.Context, status int, data interface{}, message string) {
	ctx.JSON(status, Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func Fail(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, Response{
		Success: false,
		Error: &ErrorInfo{
			Message: message,
		},
	})
}