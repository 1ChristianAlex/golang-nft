package auth

import (
	"errors"
	pwdhashing "khrix/golang-nft/src/lib/pwdHashing"
	"khrix/golang-nft/src/modules/user"
)

func LoginUser(loginUser AuthLoginInputDto) user.User {

	currentUser := user.FindUserByEmail(loginUser.Email)

	isCurrentUser := pwdhashing.ComparePasswords(currentUser.Password, loginUser.Password)

	defer func() {
		if r := recover(); r != nil || !isCurrentUser {
			panic(errors.New("wrog email or password"))
		}
	}()

	return currentUser
}
