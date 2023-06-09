package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/gabrieldebem/clickup/packages/Models"
)

type ClickUpClient struct {
    BaseUrl string
    Token string
    SpaceId string
    TeamId string
    FolderId string
    ListId string
}

func (c ClickUpClient) GetAuthorizadedUser() models.UserResponse {
  req, _ := http.NewRequest("GET", c.BaseUrl + "/v2/user", nil)
  req.Header.Add("Authorization", c.Token)
  res, _ := http.DefaultClient.Do(req)
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  var users models.UserResponse

  err := json.Unmarshal(body, &users)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(users.User.Username)
  fmt.Println(users.User.Id)

  return users
}

func (c ClickUpClient) GetFolders() models.FolderResponse {
  req, _ := http.NewRequest("GET", c.BaseUrl + "/v2/space/" + c.SpaceId + "/folder", nil)

  query := req.URL.Query()
  query.Add("archived", "false")
  req.URL.RawQuery = query.Encode()
  
  req.Header.Add("Authorization", c.Token)
  
  res, _ := http.DefaultClient.Do(req)
  
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  var folders models.FolderResponse
  json.Unmarshal(body, &folders)

  for _, folder := range folders.Folders {
    fmt.Println(folder.Name)
    fmt.Println(folder.Id)
  }

  return folders
}

func (c ClickUpClient) GetTeams() models.TeamResponse {
    req, _ := http.NewRequest("GET", c.BaseUrl + "/v2/team", nil)
    req.Header.Add("Authorization", c.Token)
    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    
    var teams models.TeamResponse
    err := json.Unmarshal(body, &teams)

    if err != nil {
        fmt.Println(err)
    }

    for _, team := range teams.Teams {
        fmt.Println(team.Name)
        fmt.Println(team.Id)
    }

    return teams
}

func (c ClickUpClient) GetSpaces() models.SpaceResponse {
    req, _ := http.NewRequest("GET", c.BaseUrl + "/v2/team/" + c.TeamId + "/space", nil)
    req.Header.Add("Authorization", c.Token)
    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    
    var spaces models.SpaceResponse
    err := json.Unmarshal(body, &spaces)

    if err != nil {
        fmt.Println(err)
    }

    for _, space := range spaces.Spaces {
        fmt.Println(space.Name)
        fmt.Println(space.Id)
    }
    
    return spaces
}

func (c ClickUpClient) GetLists() models.ListResponse {
    req, _ := http.NewRequest("GET", c.BaseUrl + "/v2/folder/" + c.FolderId + "/list", nil)
    req.Header.Add("Authorization", c.Token)
    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    
    var lists models.ListResponse
    err := json.Unmarshal(body, &lists)
    if err != nil {
        fmt.Println(err)
    }

    for _, list := range lists.Lists {
        fmt.Println(list.Name)
        fmt.Println(list.Id)
    }
    
    return lists
}

func (c ClickUpClient) GetTasks() models.TaskResponse {
    req, _ := http.NewRequest("GET", c.BaseUrl + "/v2/list/" + c.ListId + "/task", nil)
    req.Header.Add("Authorization", c.Token)
    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    
    var tasks models.TaskResponse
    err := json.Unmarshal(body, &tasks)
    if err != nil {
        fmt.Println(err)
    }

    for _, task := range tasks.Tasks {
        fmt.Println(task.Name)
        fmt.Println(task.Id)
    }
    
    return tasks
}

