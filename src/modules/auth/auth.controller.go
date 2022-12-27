package auth

import (
	"fmt"
	"khrix/golang-nft/src/modules/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleLogin(context *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			errorResponse := common.ErrorResponse{Message: fmt.Sprint(r.(error)), Status: http.StatusNotFound}

			context.JSON(errorResponse.Status, errorResponse)
		}
	}()

	var loginDto AuthLoginInputDto

	if err := context.BindJSON(&loginDto); err != nil {
		panic(err)
	}

	currentUser := LoginUser(loginDto)

	context.JSON(http.StatusOK, currentUser)

}

func RegisterController(app *gin.Engine) {
	NAME := "auth"

	app.POST(fmt.Sprintf("/%s/login", NAME), handleLogin)
}
