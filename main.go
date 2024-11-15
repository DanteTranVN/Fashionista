package main

import (
	"fashora-backend/config"
	"fashora-backend/controllers/auth_controller"
	"fashora-backend/controllers/user_controller"
	"fashora-backend/middlewares"
	"fashora-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.LoadConfig()

	// Kết nối database
	models.ConnectDatabase()

	// Routes
	r.POST("/auth/register", auth_controller.Register)
	r.POST("/auth/login", auth_controller.Login)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.PATCH("/user/update_user", user_controller.UpdateUser) // Route update người dùng
	}

	r.Run(":8080")
}
