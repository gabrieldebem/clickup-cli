package models

type List struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Content string `json:"content"`
    Archived bool `json:"archived"`
}

type ListResponse struct {
    Lists []List `json:"lists"`
}
