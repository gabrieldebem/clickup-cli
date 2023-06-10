package usecases

import clients "github.com/gabrieldebem/clickup/packages/Clients"

func GetUser(c clients.ClickUpClient) {
    c.GetAuthorizadedUser()
}
