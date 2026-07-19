package routes

import (
	"example.com/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.DELETE("/events/:id/registrations", cancelRegistration)
	authenticated.POST("/events/:id/registrations", registerForEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
