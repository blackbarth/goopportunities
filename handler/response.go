package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, status int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"errorCode": status,
		"message":   message,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler %s successful", op),
		"data":    data,
	})
}
