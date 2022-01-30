package prepare

import (
	"github.com/DanielTitkov/crud-api-service-sample/internal/api/handler"
	"github.com/DanielTitkov/crud-api-service-sample/internal/app"
	"github.com/DanielTitkov/crud-api-service-sample/internal/configs"
	"github.com/DanielTitkov/crud-api-service-sample/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg configs.Config, logger *logger.Logger, app *app.App) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Middleware
	r.Use(cors.Default())

	handler.SetupHandler(r, cfg, logger, app)
	return r
}
