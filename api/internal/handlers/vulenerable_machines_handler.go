// Vulenrable machines route handler
package handlers

import (
	"log"
	"viralvault/internal/repositories"

	"github.com/gin-gonic/gin"
)

// Define VulenerableMachineHandler struct
type VulenerableMachineHandler struct {
	vulenrableMachineRepo repositories.VulenerableMachineRepository
}

// Initialize VulenerableMachineHandler
func NewVulenerableMachineHandler(vulenrableMachineRepo repositories.VulenerableMachineRepository) *VulenerableMachineHandler {
	return &VulenerableMachineHandler{vulenrableMachineRepo: vulenrableMachineRepo}
}

// GetVulnerableMachines returns a list of all vulnerable machines
func (h *VulenerableMachineHandler) GetVulnerableMachines(c *gin.Context) {
	machines, err := h.vulenrableMachineRepo.GetVulnerableMachines()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, machines)
}

// GetVulnerableMachine returns a specific vulnerable machine
func (h *VulenerableMachineHandler) GetVulnerableMachine(c *gin.Context) {
	id := c.Param("id")
	machine, err := h.vulenrableMachineRepo.GetVulnerableMachine(id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, machine)
}
