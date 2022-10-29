package app

import (
	"khrix/golang-nft/src/modules/hello"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	app := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, HelloResponse{Message: "Opa eai"})
	// })

	hello.RegisterHelloController(app)

	app.Run(":" + ApiPort)
}
