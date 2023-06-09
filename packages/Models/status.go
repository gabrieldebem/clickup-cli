package models

type Status struct {
    Status string `json:"status"`
    Color string `json:"color"`
    Orderindex int `json:"orderindex"`
    Type string `json:"type"`
}
