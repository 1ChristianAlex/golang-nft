package user

type UserOutputDto struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Role     int32  `json:"role"`
}

type UserInputDto struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     int32  `json:"role" validate:"required"`
}

func (userInput UserOutputDto) MapUserOutputDto(userItem User) UserOutputDto {
	return UserOutputDto{Name: userItem.Name, LastName: userItem.LastName, Email: userItem.Email, Role: userItem.Role}
}
