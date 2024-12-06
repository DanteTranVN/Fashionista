package auth_service

import (
	"errors"
	"fashora-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(c *gin.Context) (models.Users, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return models.Users{}, errors.New("User not authenticated")
	}

	user, ok := userInterface.(models.Users)
	if !ok {
		return models.Users{}, errors.New("Invalid user type")
	}

	return user, nil
}
