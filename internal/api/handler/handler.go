package handler

import (
	"github.com/DanielTitkov/crud-api-service-sample/internal/app"
	"github.com/DanielTitkov/crud-api-service-sample/internal/configs"
	"github.com/DanielTitkov/crud-api-service-sample/internal/logger"
	"github.com/labstack/echo"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.POST("/createPizza", h.CreateUserHandler)
	v1.POST("/getPizza", h.GetTokenHandler)
	v1.POST("/updatePizza", h.SearchHandler)
	v1.POST("/deletePizza", h.GetPublicUserHandler)
}
