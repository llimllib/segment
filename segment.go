// Package segment provides an implementation of Norvig's recursive word
// segmenter given in http://norvig.com/ngrams/ch14.pdf
package segment

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/calmh/lfucache"
)

func getProbs(filename string) map[string]float64 {
	//just read the whole stupid file into memory
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read file", filename)
		os.Exit(1)
	}

	//split the file into words
	s := regexp.MustCompile(`\s`).Split(string(content), -1)

	//increment the counter by inc every time we encounter a word
	inc := 1.0 / float64(len(s))

	wordprobs := make(map[string]float64)

	for _, word := range s {
		word = strings.ToLower(strings.Trim(word, ",-!;:\"?."))
		_, ok := wordprobs[word]
		if ok {
			wordprobs[word] += inc
		} else {
			wordprobs[word] = inc
		}
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
func MakeWordProb(filename string) func(string) float64 {
	wordprobs := getProbs(filename)

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

// CacheSize is the maximum number of keys segment will keep in its cache
var CacheSize = 65536
var seen *lfucache.Cache

// Segment a string. Return the highest-scoring segmentation of that string
// given the word probability function wordprob.
func Segment(text string, wordprob func(string) float64) []string {
	if seen == nil {
		seen = lfucache.Create(CacheSize)
	}

	if len(text) == 0 {
		return []string{}
	}

	res, ok := seen.Access(text)
	if ok {
		return res.([]string)
	}

	candidates := make([][]string, 0)
	for _, sp := range splits(text) {
		candidates = append(candidates, append([]string{sp.Head}, Segment(sp.Tail, wordprob)...))
	}

	max := maxPword(candidates, wordprob)
	seen.Insert(text, max)

	return max
}
