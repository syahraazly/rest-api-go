package main

import (
	"rest-api/database"
	"rest-api/models"
	"rest-api/router"
	"log"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	r := router.SetupRouter()
	r.Run(":8080") 
}
