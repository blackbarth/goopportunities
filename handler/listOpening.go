package handler

import (
	"net/http"

	"github.com/blackbarth/goopportunities.git/schemas"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to list openings")
		return
	}
	
	sendSuccess(ctx, "list-openings", openings)

}