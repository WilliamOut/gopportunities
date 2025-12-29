package handler

import (
	"net/http"

	"github.com/WilliamOut/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary List openings
// @Tags Openings
// @Router /openings [get]
func ShowOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro ao listar vagas")
		return
	}

	sendSuccess(ctx, "vagas listadas", openings)
}
