package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudysacosta/go-and-mysql-dockerizing/api/models"
)

// RolesHandler is a Gin handler function that handles requests to fetch all roles.
// It takes a pointer to a role model (models.RolModel) as an argument.
// It returns a Gin handler function that handles requests related to fetching all roles.
func RolesHandler(rolModel *models.RolModel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Call List() method of RolModel
		roles, err := rolModel.List()

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to fetch roles"})
			return
		}

		ctx.JSON(200, gin.H{
			"roles": roles,
		})
	}
}

// RolHandler is a Gin handler function that handles requests related to rol.
// It takes a pointer to a role model (models.RolModel) as an argument.
// It returns a Gin handler function that handles requests related to rol.
func RolHandler(rolModel *models.RolModel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract the ID from the request parameter and convert it to an integer.
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(500, gin.H{"error": "id not valid"})
			return
		}

		rol, err := rolModel.Get(id)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to find the rol"})
			return
		}

		ctx.JSON(200, gin.H{
			"rol": rol,
		})
	}
}

func CreateHandler(rolModel *models.RolModel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rol models.Rol

		if err := ctx.ShouldBindJSON(&rol); err != nil {
			ctx.JSON(400, gin.H{"error": "invalide json data"})
			return
		}

		// Call the Create method of RolModel to create the rol
		resurt, err := rolModel.Create(&rol)

		id, _ := resurt.LastInsertId()

		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to create rol"})
			return
		}

		rol.ID = uint(id)

		// Return a success message
		ctx.JSON(201, gin.H{"message": "rol created successfully", "rol": rol})

	}
}
