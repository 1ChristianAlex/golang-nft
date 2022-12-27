package app

import (
	"fmt"
	"khrix/golang-nft/src/config"
	"khrix/golang-nft/src/modules/auth"
	"khrix/golang-nft/src/modules/common"
	"khrix/golang-nft/src/modules/hello"
	"khrix/golang-nft/src/modules/user"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	config.PrepareViper()
	app := gin.Default()

	teste := []common.RegisterController{hello.RegisterController, user.RegisterController, auth.RegisterController}

	for _, register := range teste {
		register(app)
	}

	apiUrl := fmt.Sprintf("127.0.0.1:%s", config.GetEnv("apiPort"))

	app.Run(apiUrl)
}
