package days

// Day 2 involves counting letters in strings in a litst
// If a string has at 2 of the same letter it's counted
// If a string has 3 of the same letter it's counted
// The totals of the 2-count and 3-count are multiplied together for a checksum

import (
	"bufio"
	"tools"
)

func DayTwo(scanner *bufio.Scanner) int {
	twocount, threecount := 0, 0

	for scanner.Scan() {
		twos, threes := countLetters(scanner.Text())
		if twos {
			twocount++
		}
		if threes {
			threecount++
		}
	}

	return twocount * threecount
}

func DayTwoPartTwo(scanner *bufio.Scanner) string {
	words := tools.ReadIntoStringArray(scanner)
	length := len(words)

	for i := 0; i < length; i++ {
		firstword := words[i]
		for j := i + 1; j < length; j++ {
			secondword := words[j]
			if edits, pos := editDistance(firstword, secondword); edits == 1 {
				finalstring := secondword[0:pos] + secondword[pos +1:]
				return finalstring
			}
		}
	}

	return ""
}

func countLetters(str string) (twos bool, threes bool) {
	if str == ""{
		return false, false
	}

	var twocount, threecount int = 0, 0
	letters := make(map[rune]int)

	// ababac
	for _, char := range str {
		letters[char]++
	}

	for _, count := range letters {
		if count == 2 {
			twocount++
		} else if count == 3 {
			threecount++
		}
	}

	if twocount > 0 {
		twos = true
	}

	if threecount > 0 {
		threes = true
	}

	return twos, threes
}

// editDistance finds the number of edits between two strings
// NOTE: this implementation assumes that all strings are the same length
func editDistance(str1 string, str2 string) (numedits int, position int) {
	var count, pos int = 0, -1
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			pos = i
			count++
		}
	}
	return count, pos
}