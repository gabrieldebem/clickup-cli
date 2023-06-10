package usecases

import clients "github.com/gabrieldebem/clickup/packages/Clients"

func UpdateTicket(c clients.ClickUpClient, ticketId string, status string) (resp string) {
	task := c.UpdateTask(ticketId, status)

	resp = "ID: " + task.Id + "\n" + "Name: " + task.Name + "\n" + "Description: " + task.TextContent + "\n" + "Status: " + task.Status.Status + "\n"
	return
}
