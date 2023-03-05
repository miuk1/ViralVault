package main

import (
	"log"

	"viralvault/internal/routes"
	"viralvault/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//Connect to the database
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=myusername dbname=mydatabase password=mypassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.VulnerableMachine{})
	db.AutoMigrate(&models.Vulnerability{})

	//Create a new router using the Gin web framework
	r := gin.Default()

	//Setup routes
	routes.SetupRoutes(db, r)

	//Start the server
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
