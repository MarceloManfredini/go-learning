package main

import (
	"log"

	"task-api/database"
	"task-api/handlers"
	"task-api/models"
	"task-api/repository"
	"task-api/routes"
	"task-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	// auto migrate
	if err := database.DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	repo := repository.NewTaskRepository()
	svc := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	r := gin.Default()
	routes.Register(r, handler)

	log.Println("Starting server :8080")
	r.Run(":8080")
}
