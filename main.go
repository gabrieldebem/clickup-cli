package main

import (
	"fmt"
	"os"

	clients "github.com/gabrieldebem/clickup/packages/Clients"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    token := os.Getenv("CLICKUP_TOKEN")
    baseUrl := os.Getenv("CLICKUP_BASE_URL")
    spaceId := os.Getenv("CLICKUP_SPACE_ID")
    teamId := os.Getenv("CLICKUP_TEAM_ID")
    folderId := os.Getenv("CLICKUP_FOLDER_ID")
    listId := os.Getenv("CLICKUP_LIST_ID")

    client := clients.ClickUpClient{
        BaseUrl: baseUrl,
        Token: token,
        SpaceId: spaceId,
        TeamId: teamId,
        FolderId: folderId,
        ListId: listId,
    }

    client.GetAuthorizadedUser()
}
