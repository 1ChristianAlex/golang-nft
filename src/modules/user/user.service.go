package user

import "errors"

type UserConnection struct{}

type UserService interface {
	getAllUsers() []User
	getUserById(id int32) User
}

func GetAllUsers() []User {
	return []User{
		{Name: "Christian", LastName: "Alexsander", Email: "christianalexbh@hotmail.com", Role: 1, Password: "$2a$10$Q9GcqDbfqlYPw8DOoa3KFuUc/gDfAkXNThdJXHitk8rpePXtMd3va", Id: 1},
		{Name: "Pedro", LastName: "Anjos", Email: "pedro@hotmail.com", Role: 1, Password: "algumasenha", Id: 2},
	}
}

func GetUserById(id int32) User {
	users := GetAllUsers()

	var userFinded User

	for _, users := range users {
		if users.Id == id {
			userFinded = users
		}
	}

	return userFinded
}

func FindUserByEmail(email string) User {
	users := GetAllUsers()

	var userFinded User

	for _, users := range users {
		if users.Email == email {
			userFinded = users
		}
	}

	if userFinded.Id == 0 {
		panic(errors.New("User not found"))
	}

	return userFinded
}
