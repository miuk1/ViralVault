// Go script to test the API
package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"viralvault/internal/routes"
	"viralvault/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupTestDB sets up a mock database for testing
func setupTestDB() (*gorm.DB, error) {
	// load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file for testing")
	}

	// get test database environment variables
	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	// create the database connection string
	dbURI := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	// create a new database instance for testing
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		return nil, err
	}

	// migrate the database schema
	db.AutoMigrate(&models.VulnerableMachine{})
	db.AutoMigrate(&models.Vulnerability{})
	db.AutoMigrate(&models.User{})

	return db, nil
}

// TestGetVulnerableMachines tests the GET /api/vulnerable-machines endpoint
func TestGetVulnerableMachines(t *testing.T) {
	//Set the gin to test mode and setup a mock database
	gin.SetMode(gin.TestMode)
	db, err := setupTestDB()

	// Check for database connection error
	if err != nil {
		t.Fatal("Error connecting to database: ", err)
	}

	//Initialize a new router
	r := gin.Default()

	//Setup routes
	routes.SetupRoutes(db, r)

	//Create a new request to the /api/vulnerable-machines endpoint
	req, _ := http.NewRequest("GET", "/api/vulnerable-machines", nil)
	if req == nil {
		t.Fatal("Request was nil")
	}

	//Create a new response recorder
	rr := httptest.NewRecorder()

	//Serve the request
	r.ServeHTTP(rr, req)

	//Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	//Check the response body
	var machines []models.VulnerableMachine
	errJsonRes := json.Unmarshal([]byte(rr.Body.Bytes()), &machines)
	if errJsonRes != nil {
		t.Fatal("Error getting Json response: ", errJsonRes)
	}

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)
	t.Log("Response status code: ", rr.Code)

	// Assert that the response body is not empty
	assert.NotEmpty(t, rr.Body.String())
	t.Log("Response body: ", rr.Body.String())

	// Assert that the response body is a valid JSON
	var response []models.VulnerableMachine
	if err := json.Unmarshal([]byte(rr.Body.Bytes()), &response); err != nil {
		t.Fatal("Response body is not a valid JSON. Could not unmarshal JSON:", err)
	}

}
