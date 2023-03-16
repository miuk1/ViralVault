// Go script to test the API
package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"viralvault/internal/routes"
	"viralvault/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

// setupTestDB sets up a mock database for testing
func setupTestDB() (*gorm.DB, error) {
	// create a new database instance for testing
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=myusername dbname=viralvault_test password=mypassword sslmode=disable")

	if err != nil {
		return nil, err
	}

	// migrate the database schema
	db.AutoMigrate(&models.VulnerableMachine{})
	db.AutoMigrate(&models.Vulnerability{})

	return db, nil
}

// TestGetVulnerableMachines tests the GET /api/vulnerable-machines endpoint
func TestGetVulnerableMachines(t *testing.T) {
	//Set the gin to test mode and setup a mock database
	gin.SetMode(gin.TestMode)
	db, _ := setupTestDB()

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
	err := json.Unmarshal([]byte(rr.Body.Bytes()), &machines)
	if err != nil {
		t.Fatal(err)
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
