// repository for the vulnerable machines database functions
package repositories

import (
	"viralvault/models"

	"github.com/jinzhu/gorm"
)

// Define interface for VulenerableMachineRepository
type VulenerableMachineRepository interface {
	GetVulnerableMachines() ([]models.VulnerableMachine, error)
	GetVulnerableMachine(id string) (models.VulnerableMachine, error)
}

// Define VulenerableMachineRepository struct
type vulenerableMachineRepository struct {
	db *gorm.DB
}

// Initialize VulenerableMachineRepository
func NewVulenerableMachineRepository(db *gorm.DB) VulenerableMachineRepository {
	return &vulenerableMachineRepository{db: db}
}

// GetVulnerableMachines returns a list of all vulnerable machines
func (r *vulenerableMachineRepository) GetVulnerableMachines() ([]models.VulnerableMachine, error) {
	var machines []models.VulnerableMachine
	if err := r.db.Preload("Vulnerabilities").Find(&machines).Error; err != nil {
		return nil, err
	}
	return machines, nil
}

// Get specific vulnerable machine
func (r *vulenerableMachineRepository) GetVulnerableMachine(id string) (models.VulnerableMachine, error) {
	var machine models.VulnerableMachine
	if err := r.db.Preload("Vulnerabilities").First(&machine, id).Error; err != nil {
		return machine, err
	}
	return machine, nil
}
