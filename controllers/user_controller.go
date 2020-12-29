package controllers

import (
	"github.com/Kazuchanfl/GoWebApp/services"
)

/*
UserResponse でコントローラのレスポンスを構造化
*/
type UserResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/*
GetAllUsers でユーザーを全取得
*/
func GetAllUsers() []UserResponse {
	users := services.GetAllUsers()
	response := make([]UserResponse, 0, len(users))
	for _, v := range users {
		response = append(response, UserResponse{
			Name: v.Name,
			Age:  v.Age,
		})
	}
	return response
}
