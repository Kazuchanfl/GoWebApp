package controllers

/*
UserResponse でコントローラのレスポンスを構造化
*/
type UserResponse struct {
	Name string `json:"name"`
}

/*
GetAllUsers でユーザーを全取得
*/
func GetAllUsers() []UserResponse {
	return []UserResponse{
		UserResponse{
			Name: "Kazuaki",
		},
	}
}
