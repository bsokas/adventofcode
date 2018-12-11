package days

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	var fakedate string = "[1518-09-22]"
	var year int = 1518
	var month int = 9
	var day int = 22

	actual := parseDate(fakedate)
	if actual[0] != year && actual[1] != month && actual[2] != day {
		t.Fatal("Response: ", actual)
	}
	
}