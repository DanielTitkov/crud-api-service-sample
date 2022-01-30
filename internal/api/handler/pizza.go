package handler

import (
	"net/http"

	"github.com/DanielTitkov/crud-api-service-sample/internal/api/model"
	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPizzaByIDHandler(c *gin.Context) {
	var request model.GetPizzaByIDRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to get pizza",
		})
		return
	}

	pizza, err := h.app.GetPizzaByID(c, request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to get pizza",
		})
		return
	}

	c.JSON(http.StatusOK, model.GetPizzaByIDResponse{Pizza: domainToModelPizza(pizza)})
}

func (h *Handler) DeletePizzaByIDHandler(c *gin.Context) {
	var request model.DeletePizzaByIDRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to delete pizza",
		})
		return
	}

	err := h.app.DeletePizzaByID(c, request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to delete pizza",
		})
		return
	}

	c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "pizza deleted",
	})
}

func (h *Handler) UpdatePizzaHandler(c *gin.Context) {
	var request model.UpdatePizzaRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to update pizza",
		})
		return
	}

	err := h.app.UpdatePizza(c, &domain.Pizza{
		ID:          request.ID,
		Description: request.Description,
		Price:       request.Price,
		Dough:       request.Dough,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to update pizza",
		})
		return
	}

	c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "pizza updated",
	})
}

func (h *Handler) CreatePizzaHandler(c *gin.Context) {
	var request model.CreatePizzaRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to create pizza",
		})
		return
	}

	err := h.app.CreatePizza(c, &domain.Pizza{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Dough:       request.Dough,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to create pizza",
		})
		return
	}

	c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "pizza created",
	})
}

func (h *Handler) GetPizzasHandler(c *gin.Context) {
	pizzas, err := h.app.GetPizzas(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   err.Error(),
			Message: "failed to retrieve pizzas",
		})
		return
	}

	var response model.GetPizzasResponse
	for _, p := range pizzas {
		response.Pizzas = append(response.Pizzas, domainToModelPizza(p))
	}

	c.JSON(http.StatusOK, response)
}

func domainToModelPizza(p *domain.Pizza) model.Pizza {
	return model.Pizza{
		ID:          p.ID,
		Title:       p.Title,
		Dough:       p.Dough,
		Description: p.Description,
		Price:       p.Price,
	}
}
