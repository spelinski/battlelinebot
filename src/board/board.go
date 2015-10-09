package board

type Card struct {
	Color  string
	Number int
}

type Flag struct {
	Claimer string
	North   []Card
	South   []Card
}

type Board struct {
	Flags [9]Flag
}
