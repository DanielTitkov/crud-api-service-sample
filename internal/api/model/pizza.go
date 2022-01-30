package model

type (
	// Pizza is a common model to use in various api methods
	Pizza struct {
		ID          int    `json:"id"`
		Title       string `json:"title" binding:"required"`
		Price       int64  `json:"price" binding:"required"`
		Description string `json:"description,omitempty"`
		Dough       string `json:"dough"`
	}
	CreatePizzaRequest struct {
		Pizza
	}
	GetPizzaByIDRequest struct {
		ID int `json:"id" binding:"required"`
	}
	GetPizzaByIDResponse struct {
		Pizza
	}
	GetPizzasResponse struct {
		Pizzas []Pizza `json:"pizzas"`
	}
	UpdatePizzaRequest struct {
		ID          int    `json:"id" binding:"required"`
		Price       int64  `json:"price"`
		Description string `json:"description"`
		Dough       string `json:"dough"`
	}
	DeletePizzaByIDRequest struct {
		ID int `json:"id" binding:"required"`
	}
)
