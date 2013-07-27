A Go language implementation of the Norvig segmenter, given in [this pdf](http://norvig.com/ngrams/ch14.pdf).

Licensed WTFPL; please use this code in any way you would like.

### func MakeWordProb(filename string) func(string) float64

MakeWordProb makes a word probability function from a file.

You can create your own word probability function if you want, this just provides a default implementation. The word probability function should take any word as an argument and return a float64 0 &lt;= x &lt;= 1

### func Segment(text string, wordprob func(string) float64) []string

Segment a string. Return the highest-scoring segmentation of that string given the word probability function wordprob.

# Example

```go
package main

import (
	"segment"
	"fmt"
)

func main() {
	wordp := segment.MakeWordProb("mobydick.txt")
	fmt.Println(segment.Segment("thereareshortpeopleeverywhere", wordp))
	// Output:
	// [there are short people everywhere]
}
