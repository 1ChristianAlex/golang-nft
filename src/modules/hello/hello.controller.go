package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHelloController(context *gin.Context) {

	context.JSON(http.StatusOK, HelloResponse{Message: "Opa eai"})

}

func RegisterController(app *gin.Engine) {
	app.GET("/", getHelloController)
}
