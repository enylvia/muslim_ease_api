package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseWithMeta[T any] struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    PaginatedResponse[T] `json:"data"`
}
type PaginatedResponse[T any] struct {
	Count int64 `json:"count"`
	Data  []T   `json:"data"`
}
type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Success[T any](c *gin.Context, data T, msg string) {
	c.JSON(http.StatusOK, Response[T]{
		Status:  "success",
		Message: msg,
		Data:    data,
	})
}
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, Response[any]{
		Status:  "error",
		Message: msg,
		Data:    nil,
	})
}
func PaginatedSuccess[T any](c *gin.Context, count int64, data []T, msg string) {
	c.JSON(http.StatusOK, ResponseWithMeta[T]{
		Status:  "success",
		Message: msg,
		Data: PaginatedResponse[T]{
			Count: count,
			Data:  data,
		},
	})
}
