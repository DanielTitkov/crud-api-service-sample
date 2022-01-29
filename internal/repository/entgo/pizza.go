package entgo

import (
	"context"

	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent/pizza"

	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"

	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreatePizza(ctx context.Context, p *domain.Pizza) (*domain.Pizza, error) {
	pizza, err := r.client.Pizza.
		Create().
		SetTitle(p.Title).
		SetPrice(p.Price).
		SetDescription(p.Description).
		SetDough(pizza.Dough(p.Dough)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainPizza(pizza), nil
}

func (r *EntgoRepository) FilterPizza(ctx context.Context, args *domain.FilterPizzaArgs) ([]*domain.Pizza, error) {
	// TODO
	return nil, nil
}

func (r *EntgoRepository) UpdatePizza(ctx context.Context, p *domain.Pizza) (*domain.Pizza, error) {
	storedPizza, err := r.client.Pizza.
		Query().
		Where(pizza.IDEQ(p.ID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	storedPizza, err = storedPizza.Update().
		SetPrice(p.Price).
		SetDescription(p.Description).
		SetDough(pizza.Dough(p.Dough)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainPizza(storedPizza), nil
}

func (r *EntgoRepository) GetPizzaByID(ctx context.Context, id int) (*domain.Pizza, error) {
	pizza, err := r.client.Pizza.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return entToDomainPizza(pizza), nil
}

func (r *EntgoRepository) DeletePizzaByID(ctx context.Context, id int) error {
	return r.client.Pizza.DeleteOneID(id).Exec(ctx)
}

func entToDomainPizza(p *ent.Pizza) *domain.Pizza {
	return &domain.Pizza{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Dough:       p.Dough.String(),
		Price:       p.Price,
	}
}
