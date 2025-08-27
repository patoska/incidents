package router

import (
	"os"
	"time"
	"blazestack/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var (
	r *gin.Engine
)

func SetupRouter() *gin.Engine {
	r := gin.Default()	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")}, // Specify allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/incidents", controllers.ListIncidents)
	r.POST("api/incidents", controllers.CreateIncident)

	return r		
}
