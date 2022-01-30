package app

import (
	"context"

	"github.com/DanielTitkov/crud-api-service-sample/internal/configs"
	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"
	"github.com/DanielTitkov/crud-api-service-sample/internal/logger"
)

type (
	// App combines services and holds top-level business logic
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		repo   Repository
	}
	// Repository has methods for storing and retrieving data
	Repository interface {
		// pizza
		CreatePizza(context.Context, *domain.Pizza) (*domain.Pizza, error)
		UpdatePizza(context.Context, *domain.Pizza) (*domain.Pizza, error)
		GetPizzas(context.Context) ([]*domain.Pizza, error)
		GetPizzaByID(context.Context, int) (*domain.Pizza, error) // this may be as well deprecated in favor of "FilterPizza" method
		DeletePizzaByID(context.Context, int) error
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) (*App, error) {
	app := App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}

	// init app jobs, caches and preload data (if any)

	return &app, nil
}
