package handler

import (
	"fmt"
	"net/http"

	"github.com/blackbarth/goopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)


// @BasePath /api/v1

// @Summary Delete an opening
// @Description Delete an opening by its ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// Find the opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening id %s not found", id))
		return
	}
	// Delete the opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Failed to delete opening id %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)


}