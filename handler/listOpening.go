package handler

import (
	"net/http"

	"github.com/blackbarth/goopportunities.git/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List all openings
// @Description Retrieve a list of all openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to list openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)

}