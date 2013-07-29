// Package segment provides an implementation of Norvig's recursive word
// segmenter given in http://norvig.com/ngrams/ch14.pdf
package segment

import (
	"bufio"
	"io"
	"math"
	"strings"
)

func getProbs(reader io.Reader) map[string]float64 {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	wordprobs := make(map[string]float64)

	var l float64 = 0

	for scanner.Scan() {
		word := strings.Trim(scanner.Text(), ",-!;:\"?.")
		_, ok := wordprobs[word]
		if ok {
			wordprobs[word] += 1
		} else {
			wordprobs[word] = 1
		}

		l++
	}

	// normalize by # of words
	for word, val := range wordprobs {
		wordprobs[word] = val/l
	}

	return wordprobs
}

// guessProb returns the probability for a word that's not in our
// corpus. It favors short strings. 10/(n * 10**l)
func guessProb(word string, n int) float64 {
	l := float64(len(word))
	return 10 / (float64(n) * math.Pow(10, l))
}

// MakeWordProb makes a word probability function from a file.
//
// You can create your own word probability function if you want, this
// just provides a default implementation. The word probability function
// should take any word as an argument and return a float64 0 <= x <= 1
func MakeWordProb(reader io.Reader) func(string) float64 {
	wordprobs := getProbs(reader)

	return func(word string) float64 {
		score, ok := wordprobs[word]
		if ok {
			return score
		}
		return guessProb(word, len(wordprobs))
	}
}

// Given an array of candidate segmentations, return the one with the highest
// probability given by the product of the wordprob(word) for all words
// in the string
func maxPword(words [][]string, wordprob func(string) float64) []string {
	var max []string
	maxscore := float64(-1)

	for _, candidate := range words {
		var totalscore float64 = 1
		for _, word := range candidate {
			totalscore *= wordprob(word)
		}

		if totalscore > maxscore {
			max = candidate
			maxscore = totalscore
		}
	}

	return max
}

type split struct {
	Head string
	Tail string
}

// Given a string, return all possible splits
func splits(text string) []split {
	var res []split

	for i := range text {
		res = append(res, split{text[:i+1], text[i+1:]})
	}

	return res
}

var seen = map[string][]string{}

// Segment a string. Return the highest-scoring segmentation of that string
// given the word probability function wordprob.
func Segment(text string, wordprob func(string) float64) []string {
	if len(text) == 0 {
		return []string{}
	}

	res, ok := seen[text]
	if ok {
		return res
	}

	candidates := make([][]string, 0)
	for _, sp := range splits(text) {
		candidates = append(candidates, append([]string{sp.Head}, Segment(sp.Tail, wordprob)...))
	}

	max := maxPword(candidates, wordprob)
	seen[text] = max

	return max
}
