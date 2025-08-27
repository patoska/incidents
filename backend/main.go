package main

import (
	"fmt"
	"os"

	"blazestack/router"
	"blazestack/models"
)

func main() {
	models.InitDB()
	r := router.SetupRouter()

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Print("Running server...")
	r.Run(":" + port)
}
