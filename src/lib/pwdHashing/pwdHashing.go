package pwdhashing

import (
	"fmt"
	"log"

	"khrix/golang-nft/src/config"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) string {
	passwordIncreased := increasePassword(password)

	bytePassword := []byte(passwordIncreased)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func increasePassword(password string) string {
	middleIndexPassword := int((len(password) / 2))

	extraSalt := config.GetEnv("pwdSalt")
	prefix := password[:middleIndexPassword]
	sufix := password[middleIndexPassword:]

	passwordWithSalt := fmt.Sprintf("%s%s%s", prefix, extraSalt, sufix)
	return passwordWithSalt
}

func ComparePasswords(hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(increasePassword(password))
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
