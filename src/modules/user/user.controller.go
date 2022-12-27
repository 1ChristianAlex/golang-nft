package user

import (
	"fmt"
	"net/http"
	"strconv"

	pwdHashing "khrix/golang-nft/src/lib/pwdHashing"

	"github.com/gin-gonic/gin"
)

func handleGetAllUsers(context *gin.Context) {
	users := GetAllUsers()

	var usersOutput []UserOutputDto

	for _, userItem := range users {
		usersOutput = append(usersOutput, UserOutputDto{}.MapUserOutputDto(userItem))
	}

	context.JSON(http.StatusOK, usersOutput)
}

func handleGetUser(context *gin.Context) {
	x, _ := strconv.ParseInt(context.Param("id"), 2, 32)

	userFind := GetUserById(int32(x))

	context.JSON(http.StatusOK, userFind)
}

func handleCreateUser(context *gin.Context) {
	var userInputDto UserInputDto

	if err := context.BindJSON(&userInputDto); err != nil {

		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hashed := pwdHashing.CreateHashPassword(userInputDto.Password)

	userInputDto.Password = hashed

	context.JSON(http.StatusOK, userInputDto)
}

func RegisterController(app *gin.Engine) {
	NAME := "users"

	routes := app.Group(fmt.Sprintf("/%s", NAME))

	routes.GET("/", handleGetAllUsers)
	routes.GET("/:id", handleGetUser)
	routes.POST("/", handleCreateUser)
}
