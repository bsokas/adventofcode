package days

import (
	"bufio"
	"strings"
)

// InitialString stores the original value extracted from file
var InitialString string

// DayFive provides the solution to the fifth day problem of the Advent of Code
func DayFive(partTwo bool, scanner *bufio.Scanner) int {
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	InitialString = strings.Join(lines, "")

	if partTwo {
		shortest := len(InitialString)
		alphabet := "abcdefghijklmnopqrstuvwxyz"
		for i := 0; i < len(alphabet); i++ {
			removed := removeCharacters(InitialString, alphabet[i:i+1])
			broken := BreakDown(removed)
			if len(broken) < shortest {
				shortest = len(broken)
			}
		}
		return shortest
	}

	result := BreakDown("")
	return len(result)
}

// BreakDown runs the actual analysis of the string to find reactions
func BreakDown(input string) string {
	experiment := InitialString
	if input != "" {
		experiment = input
	}

	clashes := 0
	i := 0

	for true {
		if hasReaction(string(experiment[i]), string(experiment[i+1])) {
			clashes++
			experiment = generateNewString(i, i+1, experiment)
			if i <= len(experiment)-2 {
				continue
			}
		}

		if i >= len(experiment)-2 && clashes == 0 {
			return experiment
		} else if i >= len(experiment)-2 {
			clashes = 0
			i = 0
		} else {
			i++
		}
	}

	return experiment
}

func hasReaction(char1, char2 string) bool {
	if char1 != char2 && strings.ToLower(char1) == strings.ToLower(char2) {
		return true
	}
	return false
}

func generateNewString(pos1, pos2 int, str string) string {
	return strings.Join([]string{str[0:pos1], str[pos2+1:]}, "")
}

func setInitialString(initial string) {
	InitialString = initial
}

func removeCharacters(initial string, lowercase string) string {
	uppercase := strings.ToUpper(lowercase)
	nolower := strings.Replace(initial, lowercase, "", -1)
	return strings.Replace(nolower, uppercase, "", -1)
}
