package routes

import (
	"github.com/MuhammadFelda/task-manager/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")

	logtime := controllers.NewLogtimeController()
	api.GET("/logtimes", logtime.Index)
	api.GET("/logtimes/:id", logtime.Show)
	api.POST("/logtimes", logtime.Create)
	api.DELETE("/logtimes/:id", logtime.Delete)

	skill := controllers.NewSkillController()
	api.GET("/skills", skill.Index)
	api.GET("/skills/:id", skill.Show)
	api.POST("/skills", skill.Create)
	api.DELETE("/skills/:id", skill.Delete)

	projectOwner := controllers.NewProjectOwnerController()
	api.GET("/projectOwners", projectOwner.Index)
	api.POST("/projectOwners/", projectOwner.Create)
	api.PUT("/projectOwners/:id", projectOwner.Update)
	api.DELETE("/projectOwners/:id", projectOwner.Delete)
}