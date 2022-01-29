package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/DanielTitkov/crud-api-service-sample/internal/api/model"
)

func (h *Handler) GetTagValuesHandler(c echo.Context) error {

	tagValues, err := h.app.GetTagValues()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get tag values",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tagValues)
}
