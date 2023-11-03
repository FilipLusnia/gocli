package main

import "testing"

func TestClean(t *testing.T) {

	type testCleanCase []struct {
		input    string
		expected []string
	}

	cases := testCleanCase{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO world",
			expected: []string{"hello", "world"},
		},
	}

	for _, singleCase := range cases {
		actualInput := cleanInput(singleCase.input)

		if len(actualInput) != len(singleCase.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actualInput), len(singleCase.expected))
			continue
		} else {
			for i := range actualInput {
				actualWord := actualInput[i]
				expectedWord := singleCase.expected[i]

				if actualWord != expectedWord {
					t.Errorf("%v doesn't equal %v", actualWord, expectedWord)
					continue
				}
			}
		}
	}
}
