package handler

import (
	"net/http"

	"github.com/blackbarth/goopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation failed: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {

		logger.Errorf("Failed to create opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	// response := opening.ToResponse()
	sendSuccess(ctx, "create-opening", opening)
}

