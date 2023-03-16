package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"viralvault/internal/routes"
	"viralvault/models"
)

func main() {

	//Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Get database environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//Create the database connection string and log it
	dbURI := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	//Connect to the database
	db, err := gorm.Open("postgres", dbURI)
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

	//Set trusted proxies
	r.SetTrustedProxies([]string{os.Getenv("APP_URL")})

	//Configure CORS middleware
	c := cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("APP_URL")},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	})
	r.Use(c)

	//Check for preflight requests and return 204
	r.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Origin", os.Getenv("APP_URL"))
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
