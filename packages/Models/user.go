package models

type User struct {
    Id int32 `json:"id"`
    Username string `json:"username"`
    Color string `json:"color"`
    ProfilePicture string `json:"profilePicture"`
}

type UserResponse struct {
    User User `json:"user"`
}
