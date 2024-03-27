package routes

import (
	controllers "finalproject/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	roleHandler := controllers.NewRoleController(db)

	userPath := r.Group("/user/role")

	userPath.POST("/", roleHandler.CreateRole)
	userPath.GET("/", roleHandler.GetAllRoles)
	userPath.GET("/:id", roleHandler.GetRoleByID)
	userPath.PUT("/:id", roleHandler.UpdateRole)
	userPath.DELETE("/:id", roleHandler.DeleteRole)
}
