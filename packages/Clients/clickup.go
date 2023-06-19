package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	models "github.com/gabrieldebem/clickup/packages/Models"
)

type ClickUpClient struct {
	BaseUrl  string
	Token    string
	SpaceId  string
	TeamId   string
	FolderId string
	ListId   string
	UserId   string
}

func (c ClickUpClient) baseClient(method string, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, c.BaseUrl+path, body)
	req.Header.Add("Authorization", c.Token)

	return req
}

func handleError(res *http.Response) {
	body, _ := io.ReadAll(res.Body)
    var err models.Error

	json.Unmarshal(body, &err)
    fmt.Println("Error: " + err.Err)

    return
}

func (c ClickUpClient) GetAuthorizadedUser() (users models.UserResponse) {
	req := c.baseClient("GET", "/v2/user", nil)
	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode >= 400 {
        handleError(res)

		return
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &users)

	return
}

func (c ClickUpClient) GetFolders() (folders models.FolderResponse) {
	req := c.baseClient("GET", "/v2/space/"+c.SpaceId+"/folder", nil)

	query := req.URL.Query()
	query.Add("archived", "false")
	req.URL.RawQuery = query.Encode()

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode >= 400 {
        handleError(res)

		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &folders)

	return
}

func (c ClickUpClient) GetTeams() (teams models.TeamResponse) {
	req := c.baseClient("GET", "/v2/team", nil)
	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode >= 400 {
        handleError(res)
		return
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &teams)

	return
}

func (c ClickUpClient) GetSpaces() (spaces models.SpaceResponse) {
	req := c.baseClient("GET", "/v2/team/"+c.TeamId+"/space", nil)
	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode >= 400 {
        handleError(res)
		return
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &spaces)

	return
}

func (c ClickUpClient) GetLists() (lists models.ListResponse) {
	req := c.baseClient("GET", "/v2/folder/"+c.FolderId+"/list", nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode >= 400 {
        handleError(res)
		return
	}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &lists)

	return
}

func (c ClickUpClient) GetTasks(onlyMine bool) (tasks models.TaskResponse) {
	req := c.baseClient("GET", "/v2/list/"+c.ListId+"/task", nil)

	if onlyMine {
		query := req.URL.Query()
		query.Add("assignees[]", c.UserId)
		req.URL.RawQuery = query.Encode()
	}

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode >= 400 {
        handleError(res)
		return
	}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &tasks)

	return
}

func (c ClickUpClient) FindTask(id string) (task models.Task) {
	req := c.baseClient("GET", "/v2/task/"+id, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode >= 400 {
        handleError(res)
		return
	}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &task)

	return
}

func (c ClickUpClient) UpdateTask(task string, status string) (taskResp models.Task) {
	reqBody := []byte(`{"status": "` + status + `"}`)

	req := c.baseClient("PUT", "/v2/task/"+task, bytes.NewBuffer(reqBody))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode >= 400 {
        handleError(res)
        return
	}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &taskResp)

	return
}

func GetClickUpInstance() *ClickUpClient {
	return &ClickUpClient{
		BaseUrl:  os.Getenv("CLICKUP_BASE_URL"),
		Token:    os.Getenv("CLICKUP_TOKEN"),
		SpaceId:  os.Getenv("CLICKUP_SPACE_ID"),
		ListId:   os.Getenv("CLICKUP_LIST_ID"),
		FolderId: os.Getenv("CLICKUP_FOLDER_ID"),
		TeamId:   os.Getenv("CLICKUP_TEAM_ID"),
		UserId:   os.Getenv("CLICKUP_USER_ID"),
	}
}
