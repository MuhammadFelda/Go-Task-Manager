package main

import (
	"log"
	"os"

	"github.com/MuhammadFelda/task-manager/config"
	"github.com/MuhammadFelda/task-manager/internal/controllers"
	"github.com/MuhammadFelda/task-manager/internal/models"
	"github.com/MuhammadFelda/task-manager/internal/repository"
	"github.com/MuhammadFelda/task-manager/internal/routes"
	"github.com/MuhammadFelda/task-manager/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db := config.DB

	if err := db.AutoMigrate(
		&models.Logtime{},
		&models.Project{},
		&models.ProjectOwner{},
		&models.Skill{},
		&models.Task{},
		&models.User{},
	); err != nil {
		log.Fatalf("Automigrage Failed: %v", err)
	}

	logtimeRepo := repository.NewLogtimeRepository(db)
	logtimeService := services.NewLogtimeService(logtimeRepo)
	logtimeController := controllers.NewLogtimeController(logtimeService)

	projectOwnerRepo := repository.NewProjectOwnerRepository(db)
	projectOwnerService := services.NewProjectOwnerService(projectOwnerRepo)
	projectOwnerController := controllers.NewProjectOwnerController(projectOwnerService)

	skillRepo := repository.NewSkillRepository(db)
	skillService := services.NewSkillService(skillRepo)
	skillController := controllers.NewSkillController(skillService)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})

	routes.Register(r, logtimeController, projectOwnerController, skillController, userController)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at port: %s", port)
	r.Run(":" + port)
}