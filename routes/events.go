package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// GET -> /events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch all events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// GET -> /events/:id
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "fetch the event successfully", "event": event})

}

// POST -> /events
func createEvent(context *gin.Context) {
	// token := context.Request.Header.Get("Authorization")
	// if token == "" {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorize"})
	// 	return
	// }

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Authorize.token is wrong"})
	// 	return

	// }

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse body"})
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create a event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "eventCreated", "event": event})
}

// PUT -> /events/:id
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}
	userId := context.GetInt64("userId")

	event, err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorize to update the event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event for updating"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse body"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event Updated"})
}

// DELETE -> /events/:id
func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorize to delete the event"})
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event "})
		return
	}
	err = event.DELETE()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})

}
