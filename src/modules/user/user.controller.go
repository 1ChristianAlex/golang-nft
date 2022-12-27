package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func handleGetAllUsers(context *gin.Context) {
	users := UserConn().getAllUsers()

	var usersOutput []UserOutputDto

	for _, userItem := range users {
		usersOutput = append(usersOutput, UserOutputDto{Name: userItem.Name, LastName: userItem.LastName, Email: userItem.Email, Role: userItem.Role})
	}

	context.JSON(http.StatusOK, usersOutput)
}

func handleGetUser(context *gin.Context) {
	x, _ := strconv.ParseInt(context.Param("id"), 2, 32)
	fmt.Println(x)

	userFind := UserConn().getUserById(int32(x))

	context.JSON(http.StatusOK, userFind)
}

func handleCreateUser(context *gin.Context) {

	var userInputDto UserInputDto

	if err := context.BindJSON(&userInputDto); err != nil {
		panic(err)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(userInputDto.Password), 8)

	userInputDto.Password = string(hashed)

	context.JSON(http.StatusOK, userInputDto)

}

func RegisterController(app *gin.Engine) {
	NAME := "users"
	app.GET(fmt.Sprintf("/%s", NAME), handleGetAllUsers)
	app.GET(fmt.Sprintf("/%s/:id", NAME), handleGetUser)
	app.POST(fmt.Sprintf("/%s", NAME), handleCreateUser)
}
