package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input         string
		expectedValue []string
	}{
		{
			input: "Ahoy sailor o/",
			expectedValue: []string{
				"ahoy",
				"sailor",
				"o/",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)

		if len(actual) != len(cs.expectedValue) {
			t.Errorf("The lengths are not equal: %v != %v", len(actual), len(cs.expectedValue))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expectedValue[i]

			if actualWord != expectedWord {
				t.Errorf("The words are not equal: %v != %v", actualWord, expectedWord)
			}
		}
	}
}
