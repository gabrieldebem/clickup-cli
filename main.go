package main

import (
	views "github.com/gabrieldebem/clickup/packages/Views"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	views.RunSelect()
}
