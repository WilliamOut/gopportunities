package handler

import (
	"net/http"

	"github.com/WilliamOut/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Delete opening
// @Tags Openings
// @Param id query string true "Opening ID"
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id é obrigatório")
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "vaga não encontrada")
		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro ao deletar vaga")
		return
	}

	sendSuccess(ctx, "vaga deletada", opening)
}
