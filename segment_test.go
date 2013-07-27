package segment

import (
	"testing"
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
