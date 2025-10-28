package initializers

import (
	"jwt-authentication/models"
	"log"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error syncing database: " + err.Error())
	}
}
