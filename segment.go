package main

import (
	"fmt"
)

func maxPword(words [][]string) []string {
}

type split struct {
	Head string
	Tail string
}

// Given a string, return all possible splits
func splits(text string) []split {
	var res []split

	for i := range(text) {
		res = append(res, split{text[:i+1], text[i+1:]})
	}

	return res
}

var seen map[string][]string = map[string][]string{}

// Given a string, return the highest-scoring segmentation of that string
func segment(text string) []string {
	if len(text) == 0 { return []string{} }

	res, ok := seen[text]
	if ok {
		return res
	}

	candidates := make([][]string, 0) //how much should I allocate? Effing sucks to have to define it...
	for _, sp := range(splits(text)) {
		candidates = append(candidates, append([]string{sp.Head}, segment(sp.Tail)...))
	}

	seen[text] := maxPword(candidates)

	return []string{}
}

func main() {
	fmt.Println(segment("thereareshortpeopleeverywhere"))
}
