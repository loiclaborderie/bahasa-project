package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loiclaborderie/bahasa-project/config"
	"github.com/loiclaborderie/bahasa-project/internal/dialogue"
	"github.com/loiclaborderie/bahasa-project/internal/module"
	"github.com/loiclaborderie/bahasa-project/pkg/db"
	"github.com/loiclaborderie/bahasa-project/routes"
)

func main() {
	database := db.Init()

	// userRepo := user.NewUserRepository(database)
	// userService := user.NewUserService(userRepo)
	// userHandler := user.NewUserHandler(userService)

	moduleRepo := module.NewModuleRepository(database)
	moduleService := module.NewModuleService(moduleRepo)
	moduleHandler := module.NewModuleHandler(moduleService)

	dialogueRepo := dialogue.NewRepository(database)
	dialogueService := dialogue.NewService(dialogueRepo)
	dialogueHandler := dialogue.NewHandler(dialogueService)

	router := gin.Default()
	// corsConfig.AllowOrigins = []string{config.GetEnv("ALLOW_ORIGIN", "")}
	fmt.Println("config :", config.GetEnv("ALLOW_ORIGIN", ""))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.GetEnv("ALLOW_ORIGIN", "")}, // or set to your dev frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/admin/modules", moduleHandler.GetAllModules)
	router.GET("/modules", moduleHandler.GetVisibleModules)
	router.POST("/modules", moduleHandler.Create)
	// router.GET("/modules/:id", moduleHandler.FindByID)
	router.GET("/modules/:id", moduleHandler.FindByID)
	router.GET("/dialogue/:id/speakers", dialogueHandler.GetAvailableSpeaker)
	router.GET("/dialogue/:id", dialogueHandler.GetDialogue)
	routes.RegisterAuthRoutes(database, router)
	log.Println("Server starting on :8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
