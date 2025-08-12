package pokeapi

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (pokePage, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

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