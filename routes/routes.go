package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loiclaborderie/go-gin-template/handlers"
)

// SetupRouter configures the Gin router with all application routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Hello world example
		v1.GET("/hello", handlers.HelloHandler)
		
		// User routes
		users := v1.Group("/users")
		{
			users.GET("/", handlers.GetUsers)
			users.GET("/:id", handlers.GetUser)
			users.POST("/", handlers.CreateUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}
	}

	return router
}