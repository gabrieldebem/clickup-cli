package models

type Team struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Color string `json:"color"`
    Avatar string `json:"avatar"`
}

type TeamResponse struct {
    Teams []Team `json:"teams"`
}
