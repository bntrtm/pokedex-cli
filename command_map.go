package main

import (
	"fmt"
)

func commandMap(c *config, args []string) error {

	page, err := c.pClient.GetPokePage(c.urlNext)
	if err != nil {
		return err
	}

	c.urlNext = page.Next
	c.urlPrevious = page.Previous

	for _, result := range page.Results {
		fmt.Println(result.Name)
	}

	return nil

}

func commandMapb(c *config, args []string) error {

	if c.urlPrevious == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	page, err := c.pClient.GetPokePage(c.urlPrevious)
	if err != nil {
		return err
	}

	c.urlNext = page.Next
	c.urlPrevious = page.Previous

	for _, result := range page.Results {
		fmt.Println(result.Name)
	}

	return nil
}
