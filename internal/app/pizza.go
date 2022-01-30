package app

import (
	"context"

	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"
)

func (a *App) CreatePizza(ctx context.Context, pizza *domain.Pizza) error {
	_, err := a.repo.CreatePizza(ctx, pizza)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetPizzas(ctx context.Context) ([]*domain.Pizza, error) {
	return a.repo.GetPizzas(ctx)
}

func (a *App) GetPizzaByID(ctx context.Context, id int) (*domain.Pizza, error) {
	return a.repo.GetPizzaByID(ctx, id)
}

func (a *App) DeletePizzaByID(ctx context.Context, id int) error {
	return a.repo.DeletePizzaByID(ctx, id)
}

func (a *App) UpdatePizza(ctx context.Context, pizza *domain.Pizza) error {
	_, err := a.repo.UpdatePizza(ctx, pizza)
	if err != nil {
		return err
	}

	return nil
}
