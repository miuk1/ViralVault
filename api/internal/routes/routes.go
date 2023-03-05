// package for viralvault routes
package routes

import (
	"log"
	"viralvault/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// setup routes
func SetupRoutes(db *gorm.DB, r *gin.Engine) {
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
}
