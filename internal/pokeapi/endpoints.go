// Package pokeapi provides types relevant to PokeAPI requests
// made by the client.
package pokeapi

import "fmt"

const (
	sVerb = "/%s"

	baseURL = "https://pokeapi.co/api/v2"

	URLLocationAreas = baseURL + "/location-area"
	URLLocationArea  = URLLocationAreas + sVerb

	URLPokemonStats = baseURL + "/pokemon"
	URLPokemonStat  = URLPokemonStats + sVerb
)

func EndpointLocationArea(name string) string {
	return fmt.Sprintf(URLLocationArea, name)
}

func EndpointPokemonStat(name string) string {
	return fmt.Sprintf(URLPokemonStat, name)
}
