package domain

type (
	FilterPizzaArgs struct {
		ID     []int
		Dough  []string
		Title  []string
		Limit  int
		Offset int
		// TODO: add filtering by create/update time
	}
)
