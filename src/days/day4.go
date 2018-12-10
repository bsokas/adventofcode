package days

// Guard timestamps are written in [year-month-day hour:minute] format

import (
	"bufio"
	"strings"
	"strconv"
)

const (
	Year = 0
	Month = 1
	Day = 2
	Hour = 3
	Minute = 4
	DayFourID = 5
	Action = 6
	EntryFields = 7
)

func DayFour(partone bool, scanner *bufio.Scanner) int {
	if !partone {return 0}

	return 0
}

// ASSUMPTIONS: entries either have 6 'words' or four, 6 if it contains a guard ID
// func parseEntry (entry string) []int {
// 	values := strings.Split(entry, " ")

// 	if len(values) == 6 {
// 		//parse a full entry
// 	} else if len(values) == 4 {
// 		//parse the time entry
// 	}

// 	return make([]int, 1)
// }

// Sample entry: '[1518-09-22 23:50] Guard #2309 begins shift'
// func guardEntry(values []string) []int {
// 	if len(values) != 6 {
// 		return nil
// 	}

// 	guardEntry := make([]int, EntryFields)

// 	return nil
// }

// func getDateTime(values []string) []int {

// }

func parseDate(raw string) []int {
	initial := strings.Split(raw, "-")
	if len(raw) != 3 {
		return nil
	}

	date := make([]int, 3)
	date[0], _ = strconv.Atoi(initial[0][1:])
	date[1], _ = strconv.Atoi(initial[1])
	date[2], _ = strconv.Atoi(initial[2][:len(initial[2]) - 1])
	return date
}