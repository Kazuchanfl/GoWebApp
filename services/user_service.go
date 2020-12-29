package services

func GetAllUsers() []UserDto {
	return []UserDto{
		UserDto{
			Name: "Kazuaki",
			Age:  23,
		},
		UserDto{
			Name: "Risa",
			Age:  20,
		},
	}
}
