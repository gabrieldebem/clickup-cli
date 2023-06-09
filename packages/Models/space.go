package models

type Space struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Private bool `json:"private"`
}


type SpaceResponse struct {
    Spaces []Space `json:"spaces"`
}
