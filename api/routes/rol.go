package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rudysacosta/go-and-mysql-dockerizing/api/handlers"
	"github.com/rudysacosta/go-and-mysql-dockerizing/api/models"
)

func InitRoutes(router *gin.Engine) {
	// Create an instance of RolModel
	rolModel := &models.RolModel{}
	router.GET("/roles", handlers.RolesHandler(rolModel))
	router.GET("/roles/:id", handlers.RolHandler(rolModel))
}
