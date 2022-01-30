package handler

import (
	"github.com/DanielTitkov/crud-api-service-sample/internal/app"
	"github.com/DanielTitkov/crud-api-service-sample/internal/configs"
	"github.com/DanielTitkov/crud-api-service-sample/internal/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func SetupHandler(
	r *gin.Engine,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(r)
	return h
}

func (h *Handler) link(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("/createPizza", h.CreatePizzaHandler)
	v1.POST("/getPizzaByID", h.GetPizzaByIDHandler)
	v1.POST("/getPizzas", h.GetPizzasHandler)
	v1.POST("/updatePizza", h.UpdatePizzaHandler)
	v1.POST("/deletePizzaByID", h.DeletePizzaByIDHandler)
}
