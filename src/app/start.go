package app

import (
	"fmt"
	"khrix/golang-nft/src/modules/common"
	"khrix/golang-nft/src/modules/hello"
	"khrix/golang-nft/src/modules/user"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	app := gin.Default()

	teste := []common.RegisterController{hello.RegisterController, user.RegisterController}

	for _, register := range teste {
		register(app)
	}

	apiUrl := fmt.Sprintf("127.0.0.1:%s", ApiPort)

	app.Run(apiUrl)
}
