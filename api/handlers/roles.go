package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rudysacosta/go-and-mysql-dockerizing/api/models"
)

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
