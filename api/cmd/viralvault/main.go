package main

import (
	"log"

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
	db.AutoMigrate(&models.VulnerableMachine{})
	db.AutoMigrate(&models.Vulnerability{})

	//Create a new router using the Gin web framework
	r := gin.Default()

	//API endpoint that retrieves a list of all vulnerable machines
	r.GET("/api/vulnerable-machines", func(ctx *gin.Context) {
		var machines []models.VulnerableMachine
		if err := db.Preload("Vulnerabilities").Find(&machines).Error; err != nil {
			ctx.AbortWithStatus(500)
			log.Fatal(err)
		} else {
			ctx.JSON(200, machines)
		}
	})

	//API endpoint that retreives a specific vulnerable machine
	r.GET("/api/vulnerable-machines/:id", func(ctx *gin.Context) {
		var machine models.VulnerableMachine
		id := ctx.Param("id")
		if err := db.Preload("Vulnerabilities").First(&machine, id).Error; err != nil {
			ctx.AbortWithStatus(404)
		} else {
			ctx.JSON(200, machine)
		}
	})

	//Start the server
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
