package main

import (
	"fmt"
	"sort"
	"errors"
)

func findMax(input []interface{}, comparator func(a interface{}, b interface{}) bool) (interface{}, error) {
	if (len(input) == 0) {
		return nil, errors.New("Input length should be more than 0")
	}

	sort.Slice(input, func(i, j int) bool {
		return comparator(input[i], input[j])
	})

	return input[0], nil
}

type WordsCounter struct {
	Word  string
	Count int
}

func main() {
	array := []interface{}{
		WordsCounter{Word: "asdf", Count: 5},
		WordsCounter{Word: "asdfgh", Count: 1},
		WordsCounter{Word: "sdgfhgfgdfsd", Count: 7},
		WordsCounter{Word: "asdfghf", Count: 4},
	}

	comparator := func(a, b interface{}) bool {
		return a.(WordsCounter).Count > b.(WordsCounter).Count
	}
	
	maxValue, error := findMax(array, comparator)
	
	if (error != nil) {
		fmt.Printf("%v", error)
	} else {
		fmt.Println(maxValue)
	}
}
