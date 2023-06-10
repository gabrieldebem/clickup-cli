package clients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

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

	return task
}

func (c ClickUpClient) UpdateTask(task string, status string) (taskResp models.Task) {
	reqBody := []byte(`{"status": "` + status + `"}`)

	req := c.baseClient("PUT", "/v2/task/"+task, bytes.NewBuffer(reqBody))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		handleError(res)
	}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &taskResp)

	return
}
