package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rudysacosta/go-and-mysql-dockerizing/api/handlers"
	"github.com/rudysacosta/go-and-mysql-dockerizing/api/models"
)

func InitRoutes(router *gin.Engine) {
	// Create an instance of RolModel
	rolModel := &models.RolModel{}
	router.GET("/roles", handlers.RolesHandler(rolModel))
	router.GET("/roles/:id", handlers.RolHandler(rolModel))
	router.POST("/roles", handlers.CreateHandler(rolModel))

	router.PUT("/roles/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		name := ctx.PostForm("name")
		description := ctx.PostForm("description")

		fmt.Printf("\n\n\n\n")
		fmt.Printf("Id : %s \n", id)
		fmt.Printf("Name from form: %s \n", name)
		fmt.Printf("description from form: %s \n", description)
		fmt.Printf("\n\n\n\n")
	})

	router.DELETE("/roles/:id/delete", func(ctx *gin.Context) {
		id := ctx.Param("id")

		fmt.Printf("\n\n\n\n")
		fmt.Printf("Id : %s \n", id)
		fmt.Printf("\n\n\n\n")
	})
}
