package prepare

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/DanielTitkov/crud-api-service-sample/internal/api/handler"
	"github.com/DanielTitkov/crud-api-service-sample/internal/app"
	"github.com/DanielTitkov/crud-api-service-sample/internal/configs"
	"github.com/DanielTitkov/crud-api-service-sample/internal/logger"
)

func NewServer(cfg configs.Config, logger *logger.Logger, app *app.App) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	if cfg.Env != "dev" {
		e.Use(middleware.Recover())
	}
	handler.NewHandler(e, cfg, logger, app)
	return e
}
