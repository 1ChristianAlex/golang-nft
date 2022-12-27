package auth

import (
	"fmt"
	"net/http"

	"khrix/golang-nft/src/modules/common"
	"khrix/golang-nft/src/modules/user"

	"github.com/gin-gonic/gin"
)

func handleLogin(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			errorResponse := common.ErrorResponse{Message: fmt.Sprint(r.(error)), Status: http.StatusNotFound}

			context.AbortWithStatusJSON(errorResponse.Status, errorResponse)
		}
	}()

	var loginDto AuthLoginInputDto

	if err := context.BindJSON(&loginDto); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse{Message: fmt.Sprint(err), Status: http.StatusBadRequest})

		return
	}

	currentUser := LoginUser(loginDto)

	context.JSON(http.StatusOK, user.UserOutputDto{}.MapUserOutputDto(currentUser))
}

func RegisterController(app *gin.Engine) {
	NAME := "auth"

	app.POST(fmt.Sprintf("/%s/login", NAME), handleLogin)
}
