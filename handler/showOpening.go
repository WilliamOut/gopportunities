package handler

import (
	"net/http"

	"github.com/WilliamOut/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Show opening
// @Tags Openings
// @Param id query string true "Opening ID"
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id é obrigatório")
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Vaga não encontrada")
		return
	}

	sendSuccess(ctx, "Vaga encontrada", opening)
}
