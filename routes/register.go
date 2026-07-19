package routes

import (
	"net/http"
	"strconv"

	"example.com/gin/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	userId := context.GetInt64("userId")

	if err := models.RegisterUserForEvent(userId, eventID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User registered for event successfully",
	})
}

func cancelRegistration(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	userId := context.GetInt64("userId")

	if err := models.CancelUserRegistration(userId, eventID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User registration canceled successfully",
	})
}
