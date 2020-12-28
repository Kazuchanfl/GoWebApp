package controllers

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
	return []UserResponse{
		UserResponse{
			Name: "Kazuaki",
			Age:  23,
		},
		UserResponse{
			Name: "Risa",
			Age:  20,
		},
	}
}
