package routes

import (
	"errors"
	"example/rest-api/models"
	"example/rest-api/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	token := context.Request.Header.Get("Authorization")

	userId, err := validateJWT(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"code":    "REAPI401",
		})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

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

	event.UserID = userId
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created!",
		"event":   event,
	})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"code":    "REAPI400",
		})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve event by Id",
			"code":    "REAPI500",
		})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event.",
			"code":    "REAPI500",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
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

	err = models.UpdateEvent(&updatedEvent)
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

func validateJWT(token string) (int64, error) {

	if token == "" {
		fmt.Println("Missing authorization token")
		return 0, errors.New("Missing authorization token")
	}

	userId, err := utils.ValidateJWTToken(token)
	if err != nil {
		fmt.Println("Invalid or expired authorization token")
		return 0, err
	}

	return userId, nil
}
