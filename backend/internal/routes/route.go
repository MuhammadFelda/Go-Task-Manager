package routes

import (
	"github.com/MuhammadFelda/task-manager/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Register(
	r *gin.Engine,
	logtimeCtrl *controllers.LogtimeController,
	projectOwnerCtrl *controllers.ProjectOwnerController,
	skillCtrl *controllers.SkillController,
	userCtrl *controllers.UserController,
) {
	api := r.Group("/api")

	api.GET("/logtimes", logtimeCtrl.Index)
	api.GET("/logtimes/:id", logtimeCtrl.Show)
	api.POST("/logtimes", logtimeCtrl.Create)
	api.DELETE("/logtimes/:id", logtimeCtrl.Delete)

	api.GET("/skills", skillCtrl.Index)
	api.GET("/skills/:id", skillCtrl.Show)
	api.POST("/skills", skillCtrl.Create)
	api.DELETE("/skills/:id", skillCtrl.Delete)

	api.GET("/projectOwners", projectOwnerCtrl.Index)
	api.POST("/projectOwners/", projectOwnerCtrl.Create)
	api.PUT("/projectOwners/:id", projectOwnerCtrl.Update)
	api.DELETE("/projectOwners/:id", projectOwnerCtrl.Delete)

	api.GET("/users", userCtrl.Index)
	api.POST("/users", userCtrl.Create)
	api.PUT("/users/:id", userCtrl.Update)
	api.DELETE("/users/:id", userCtrl.Delete)
}