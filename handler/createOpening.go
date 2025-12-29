package handler

import (
	"net/http"

	"github.com/WilliamOut/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body schemas.Opening true "Request body"
// @Success 200 {object} schemas.Opening
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := struct {
		Role     string  `json:"role"`
		Company  string  `json:"company"`
		Location string  `json:"location"`
		Remote   bool    `json:"remote"`
		Link     string  `json:"link"`
		Salary   float64 `json:"salary"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if request.Role == "" || request.Company == "" {
		sendError(ctx, http.StatusBadRequest, "cargo e empresa são obrigatórios")
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Erro ao criar vaga: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Erro ao criar vaga")
		return
	}

	sendSuccess(ctx, "Vaga criada", opening)
}
