package usecases

import (
	clients "github.com/gabrieldebem/clickup/packages/Clients"
)

func GetTickets(c clients.ClickUpClient, onlyMine bool) string {
	taskResponse := c.GetTasks(onlyMine)

	var resp string

	for _, task := range taskResponse.Tasks {
        resp += "ID: " + task.Id + "\n" + "Status: " + task.Status.Status + "\n" + "Name: " + task.Name + "\n\n"
	}

	return resp
}
