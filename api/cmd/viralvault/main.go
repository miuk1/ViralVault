package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Vulnerability struct {
	gorm.Model
	Name       string `json:"name"`
	Difficulty string `json:"difficulty"`
	MachineID  uint   `json:"machine_id"`
}

type VulnerableMachine struct {
	gorm.Model
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	OS              string          `json:"os"`
	DockerImage     string          `json:"docker_image"`
	Vulnerabilities []Vulnerability `gorm:"foreignKey:MachineID" json:"vulnerabilities"`
}

func main() {
	//Connect to the database
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=myusername dbname=mydatabase password=mypassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&VulnerableMachine{})
	db.AutoMigrate(&Vulnerability{})

	//Create a new router using the Gin web framework
	r := gin.Default()

	//API endpoint that retrieves a list of all vulnerable machines
	r.GET("/api/vulnerable-machines", func(ctx *gin.Context) {
		var machines []VulnerableMachine
		if err := db.Preload("Vulnerabilities").Find(&machines).Error; err != nil {
			ctx.AbortWithStatus(500)
			log.Fatal(err)
		} else {
			ctx.JSON(200, machines)
		}
	})

	//API endpoint that retreives a specific vulnerable machine
	r.GET("/api/vulnerable-machines/:id", func(ctx *gin.Context) {
		var machine VulnerableMachine
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
