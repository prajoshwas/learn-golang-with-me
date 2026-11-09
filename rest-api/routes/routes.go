package routes

import (
	"example/rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("events", getEvents)
	router.GET("events/:id", getEvent)

	authenticated := router.Group("/")

	authenticated.Use(middleware.Authenticate)
	authenticated.POST("events", createEvent)
	authenticated.PUT("events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.POST("events/:id/cancel", cancelRegistration)

	router.POST("signup", registerUser)
	router.POST("login", loginUser)
}
