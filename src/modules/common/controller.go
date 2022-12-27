package common

import "github.com/gin-gonic/gin"

type RegisterController func(app *gin.Engine)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
