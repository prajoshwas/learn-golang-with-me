package routes

import (
	"example/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userId")
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
			"message": "Error occurred while fetching event.",
			"code":    "REAPI500",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred while registering user for event.",
			"code":    "REAPI500",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered user for event.",
		"code":    "REAPI201",
	})
}

func cancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"code":    "REAPI400",
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found.",
			"code":    "REAPI404",
		})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		if err.Error() == "No registration found for this user and event" {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "Registration not found.",
				"code":    "REAPI404",
			})

		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error occurred while canceling registration.",
				"code":    "REAPI500",
			})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registration cancelled for event.",
		"code":    "REAPI200",
	})
}
