// model for vulnerable machine and vulnerabilities
package models

import (
	"github.com/jinzhu/gorm"
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
