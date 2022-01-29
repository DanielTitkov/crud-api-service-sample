package model

type (
	// Pizza is a common model to use in various api methods
	Pizza struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Dough       string `json:"dough"`
		Price       int64  `json:"price"`
	}
	CreatePizzaRequest struct {
		Pizza
	}
	GetPizzaByIDRequest struct {
		ID int `json:"id"`
	}
	GetPizzaByIDResponse struct {
		ID int `json:"id"`
	}
	FilterPizzaRequest struct {
		ID    []int    `json:"id"`
		Title []string `json:"title"`
		Dough []string `json:"dough"`
		Price []int64  `json:"price"` // TODO: add less-than/greater-than operators
	}
	FilterPizzaResponse struct {
		Result []Pizza `json:"result"`
	}
	UpdatePizzaByIDRequest struct {
		ID int `json:"id"`
	}
	DeletePizzaByIDRequest struct {
		ID int `json:"id"`
	}
)
