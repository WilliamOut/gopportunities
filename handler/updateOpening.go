package handler

import (
	"net/http"

	"github.com/WilliamOut/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Update opening
// @Tags Openings
// @Param id query string true "Opening ID"
// @Param request body schemas.Opening true "Request body"
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "o parâmetro ID é obrigatório")
		return
	}

	request := struct {
		Role     string `json:"role"`
		Company  string `json:"company"`
		Location string `json:"location"`
		Remote   *bool  `json:"remote"`
		Link     string `json:"link"`
		Salary   int64  `json:"salary"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "vaga não encontrada")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("erro ao atualizar vaga: %v", err)
		sendError(ctx, http.StatusInternalServerError, "erro ao salvar atualização")
		return
	}

	sendSuccess(ctx, "vaga atualizada", opening)
}
