package days

import (
	"fmt"
	"testing"
)

func TestGenerateNewString(t *testing.T) {
	tester := "morningishere"
	expected := "morninghere"

	actual := generateNewString(7, 8, tester)
	if actual != expected {
		t.Fatalf("Expected %s but function instead returned %s", expected, actual)
	}
}

func TestSmallString(t *testing.T) {
	initial := "dabAcCaCBAcCcaDA "
	setInitialString(initial)
	item := BreakDown("")

	fmt.Println(item)
}

func TestRemoveCharacters(t *testing.T) {
	initial := "dabAcCaCBAcCcaDA"
	removed := removeCharacters(initial, "d")
	broken := BreakDown(removed)

	if len(broken) != 6 {
		t.Fatalf("Was expecting length of 4 but instead got broken down string of %s", broken)
	}

}
