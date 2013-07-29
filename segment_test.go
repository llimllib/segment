package segment

import (
	"testing"
	"fmt"
	"os"
)

func sliceEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func getFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file", filename)
		os.Exit(1)
	}

	return f
}

func TestSegment(t *testing.T) {
	wordp := MakeWordProb(getFile("mobydick.txt"))
	seg := Segment("thereareshortpeopleeverywhere", wordp)

	if !sliceEq(seg, []string{"there", "are", "short", "people", "everywhere"}) {
		t.FailNow()
	}
}

func ExampleSegment() {
	wordp := MakeWordProb(getFile("mobydick.txt"))
	fmt.Println(Segment("thereareshortpeopleeverywhere", wordp))
	// Output:
	// [there are short people everywhere]
}

