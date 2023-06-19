package main

import (
	"fmt"
	"os"

	clients "github.com/gabrieldebem/clickup/packages/Clients"
	usecases "github.com/gabrieldebem/clickup/packages/UseCases"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	args := os.Args[1:]

	if len(args) > 0 {
        c := clients.GetClickUpInstance()

		switch args[0] {
		case "list":
			if len(args) > 1 && args[1] == "--only-mine" {
				fmt.Println(usecases.GetTickets(c, true))

				return
			}

			fmt.Println(usecases.GetTickets(c, false))
		case "show":
			if len(args) < 2 {
				fmt.Println("Missing ticket id")

				return
			}

			ticketId := args[1]
			fmt.Println(usecases.ShowTicket(c, ticketId))
		case "update":
			if len(args) < 3 {
				fmt.Println("Missing ticket id or new ticket status")

				return
			}

			ticketId := args[1]
			newStatus := args[2]

			fmt.Println(usecases.UpdateTicket(c, ticketId, newStatus))
		}
	}
}
