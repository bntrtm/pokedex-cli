package pokeapi

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (PokemonStat, error) {
	url := baseUrl + "/pokemon" + "/" + pokemon

	// don't make a request if a response is already cached
	if val, ok := c.cache.Get(url); ok {
		resp := PokemonStat{}
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return PokemonStat{}, err
		}

		return resp, nil
	}

	// make a request
	req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return PokemonStat{}, err
        }
        resp, err := c.httpClient.Do(req)
        if err != nil {
			return PokemonStat{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
			fmt.Println(fmt.Sprintf("HTTP error %d for request at /%s", resp.StatusCode, url))
			return PokemonStat{}, err
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
			return PokemonStat{}, err
	}

	var stats PokemonStat
	err = json.Unmarshal(jsonData, &stats)
	if err != nil {
			return PokemonStat{}, err
	}

	return stats, nil
}

func (c *Client) GetLocation(area string) (LocationArea, error) {
	url := baseUrl + "/location-area" + "/" + area

	// don't make a request if a response is already cached
	if val, ok := c.cache.Get(url); ok {
		resp := LocationArea{}
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return LocationArea{}, err
		}

		return resp, nil
	}

	// make a request
	req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return LocationArea{}, err
        }
        resp, err := c.httpClient.Do(req)
        if err != nil {
			return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
			fmt.Println(fmt.Sprintf("HTTP error %d for request at /%s", resp.StatusCode, url))
			return LocationArea{}, err
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
			return LocationArea{}, err
	}

	var locA LocationArea
	err = json.Unmarshal(jsonData, &locA)
	if err != nil {
			return LocationArea{}, err
	}

	return locA, nil
}

func (c *Client) GetPokePage(pageUrl *string) (pokePage, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// don't make a request if a response is already cached
	if val, ok := c.cache.Get(url); ok {
		resp := pokePage{}
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return pokePage{}, err
		}

		return resp, nil
	}

	// make a request
	req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return pokePage{}, err
        }
        resp, err := c.httpClient.Do(req)
        if err != nil {
			return pokePage{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
			fmt.Println(fmt.Sprintf("HTTP error %d for request at /%s", resp.StatusCode, url))
			return pokePage{}, err
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
			return pokePage{}, err
	}

	var page pokePage
	err = json.Unmarshal(jsonData, &page)
	if err != nil {
			return pokePage{}, err
	}

	return page, nil
}