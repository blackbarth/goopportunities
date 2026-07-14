package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
		// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	InitializeRouter(r)

	// Start the server on port 8080	

	r.Run(":8080")
}