package usecases

import clients "github.com/gabrieldebem/clickup/packages/Clients"

func ShowTicket(c clients.ClickUpClient, id string) (resp string) {
	ticket := c.FindTask(id)

    resp = "ID: " + ticket.Id + "\n" + "Name: " + ticket.Name + "\n" + "Description: " + ticket.TextContent + "\n" + "Status: " + ticket.Status.Status + "\n"
	return
}
