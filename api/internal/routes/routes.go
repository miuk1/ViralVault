// package for viralvault routes
package routes

import (
	"viralvault/internal/handlers"
	"viralvault/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// setup routes
func SetupRoutes(db *gorm.DB, r *gin.Engine) {
	// create a new repository instance
	vulenrableMachineRepo := repositories.NewVulenerableMachineRepository(db)

	// create a new instance of the VulenerableMachineHandler
	vulenerableMachineHandler := handlers.NewVulenerableMachineHandler(vulenrableMachineRepo)

	// API endpoint that retrieves a list of all vulnerable machines
	r.GET("/api/vulnerable-machines", vulenerableMachineHandler.GetVulnerableMachines)

	// API endpoint that retrieves a specific vulnerable machine
	r.GET("/api/vulnerable-machines/:id", vulenerableMachineHandler.GetVulnerableMachine)
}
