package user

type UserConnection struct {
}

type UserService interface {
	getAllUsers() []User
	getUserById(id int32) User
}

func (connection UserConnection) getAllUsers() []User {
	return []User{
		{Name: "Christian", LastName: "Alexsander", Email: "christianalexbh@hotmail.com", Role: 1, Password: "algumasenha", Id: 1},
		{Name: "Pedro", LastName: "Anjos", Email: "pedro@hotmail.com", Role: 1, Password: "algumasenha", Id: 2},
	}
}

func (connection UserConnection) getUserById(id int32) User {
	users := connection.getAllUsers()

	var userFinded User

	for _, users := range users {
		if users.Id == id {
			userFinded = users
		}
	}

	return userFinded
}

func UserConn() UserConnection {
	return UserConnection{}
}
