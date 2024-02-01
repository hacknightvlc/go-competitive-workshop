package cmd

type PokemonAbility struct {
	Name string
}

type PokemonType struct {
	Name string
}

type Pokemon struct {
	Name      string
	Types     []PokemonType
	Abilities []PokemonAbility
}
