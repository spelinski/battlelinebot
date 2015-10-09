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

func (b *Board)HandleFlagClaimCommand(command []string) {
	for i, claimer := range command {
		b.Flags[i].Claimer = claimer
	}
}
