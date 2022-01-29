package domain

type (
	// Pizza holds pizza data for domain logic
	Pizza struct {
		ID          int
		Title       string
		Description string
		Dough       string
		Price       int64 // price in cents
	}
)
