package models

type Task struct {
    Id string `json:"id"`
    CustomId string `json:"custom_id"`
    Name string `json:"name"`
    Status Status `json:"status"`
    TextContent string `json:"text_content"`
    Description string `json:"description"`
    DateCreated string `json:"date_created"`
    DateUpdated string `json:"date_updated"`
}

type TaskResponse struct {
    Tasks []Task `json:"tasks"`
}
