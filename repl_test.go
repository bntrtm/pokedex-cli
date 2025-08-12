package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input		string
	expected	[]string
	}{
		{
			input:		" hello world ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:		" Bulbasau r  ",
			expected:	[]string{"bulbasau", "r"},
		},
		{
			input:		"Hello, World!",
			expected:	[]string{"hello,", "world!"},
		},
		{
			input:		"Charmander Bulbasaur PIKACHU",
			expected:	[]string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:		"bLoSSom",
			expected:	[]string{"blossom"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("ERROR: input vs expected are of unequal lengths")
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("ERROR: input word is unequal to expected word")
				fmt.Println("expected: ", expectedWord)
				fmt.Println("actual: ", word)

			}
		}
	}
}