package main

import (
	"fmt"
	"os"

	clients "github.com/gabrieldebem/clickup/packages/Clients"
	"github.com/joho/godotenv"
)

func main() {
    Setup()

    token := os.Getenv("CLICKUP_TOKEN")
    baseUrl := os.Getenv("CLICKUP_BASE_URL")
    spaceId := os.Getenv("CLICKUP_SPACE_ID")
    teamId := os.Getenv("CLICKUP_TEAM_ID")
    folderId := os.Getenv("CLICKUP_FOLDER_ID")
    listId := os.Getenv("CLICKUP_LIST_ID")
    userId := os.Getenv("CLICKUP_USER_ID")

    client := clients.ClickUpClient{
        BaseUrl: baseUrl,
        Token: token,
        SpaceId: spaceId,
        TeamId: teamId,
        FolderId: folderId,
        ListId: listId,
        UserId: userId,
    }

    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("Please, provide a command")
    }

    switch args[0] {
    case "get-user":
        client.GetAuthorizadedUser()
    case "get-lists":
        client.GetLists()
        case "get-folders":
        client.GetFolders()
    case "get-tasks":
        onlyMineTickets := false
        if len(args) > 1 && args[1] == "--mine" {
            onlyMineTickets = true
        }

        client.GetTasks(onlyMineTickets)
        case "get-spaces":
        client.GetSpaces()
    case "get-teams":
        client.GetTeams()
    case "find-task":
        client.FindTask(args[1])
    case "update-task":
        if (len(args) < 3) {
            fmt.Println("Please, provide a task id and a new status")
            return
        }

        client.UpdateTask(args[1], args[2])
    default:
        fmt.Println("Command not found")
    }
}

func Setup() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
}
