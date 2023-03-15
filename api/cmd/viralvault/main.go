package main

import (
	"log"

	"viralvault/internal/routes"
	"viralvault/models"

	"github.com/gin-contrib/cors"
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

	//Set trusted origins
	r.SetTrustedProxies([]string{"127.0.0.1/8"})

	//Configure CORS middleware
	c := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	})
	r.Use(c)

	//Check for preflight requests and return 204
	r.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			ctx.AbortWithStatus(204)
		}
		ctx.Next()
	})

	//Setup routes
	routes.SetupRoutes(db, r)

	//Start the server
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
