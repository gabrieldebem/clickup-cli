package models

type Task struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Status Status `json:"status"`
    DateCreated string `json:"date_created"`
    DateUpdated string `json:"date_updated"`
}

type TaskResponse struct {
    Tasks []Task `json:"tasks"`
}
