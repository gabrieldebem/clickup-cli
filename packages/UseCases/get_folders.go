package usecases

import clients "github.com/gabrieldebem/clickup/packages/Clients"

func GetFolders(c clients.ClickUpClient) {
    c.GetFolders()
}
