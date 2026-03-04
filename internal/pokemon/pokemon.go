package pokemon

// Pokemon represents a Pokemon with its national dex ID and name.
type Pokemon struct {
	ID   int
	Name string
}

// Starters returns the three Gen 1 starter Pokemon.
func Starters() []Pokemon {
	return []Pokemon{
		{ID: 1, Name: "Bulbasaur"},
		{ID: 4, Name: "Charmander"},
		{ID: 7, Name: "Squirtle"},
	}
}
