package models

type Folder struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Hidden bool `json:"hidden"`
    TaskCount string `json:"task_count"`
}

type FolderResponse struct {
    Folders []Folder `json:"folders"`
}
