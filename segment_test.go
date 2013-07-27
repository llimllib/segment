package segment

import (
	"testing"
	"fmt"
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

func TestSegment(t *testing.T) {
	wordp := MakeWordProb("mobydick.txt")
	seg := Segment("thereareshortpeopleeverywhere", wordp)

	if !sliceEq(seg, []string{"there", "are", "short", "people", "everywhere"}) {
		t.FailNow()
	}
}

func ExampleSegment() {
	wordp := MakeWordProb("mobydick.txt")
	fmt.Println(Segment("thereareshortpeopleeverywhere", wordp))
	// Output:
	// [there are short people everywhere]
}
