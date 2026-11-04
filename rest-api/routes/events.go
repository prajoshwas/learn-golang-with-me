package routes

import (
	"example/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("events", getEvents)
	router.GET("events/:id", getEvent)
	router.POST("events", createEvent)
	router.PUT("events/:id", updateEvent)
}

func getEvents(context *gin.Context) {
	es, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events. Try again later",
			"code":    "REAPI500",
		})
		return
	}
	context.JSON(http.StatusOK, es)
}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"code":    "REAPI400",
		})
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve event by Id",
			"code":    "REAPI500",
		})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"code":    "REAPI400",
		})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event. Try again later",
			"code":    "REAPI500",
		})
		return
	}

	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created!",
		"event":   event,
	})

}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"code":    "REAPI400",
		})
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil && event == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve event by Id",
			"code":    "REAPI500",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"code":    "REAPI400",
		})
		return
	}

	updatedEvent.ID = eventId

	_, err = models.UpdateEvent(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event.",
			"code":    "REAPI500",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}
