package days

/**
	Day 3 involves finding conflicting rectangles on an area of fabric
	Input offers dimensions of [# points inward from left, # points down from top, wDayThreeIDth, height]
	Sample: [#123 @ 3,2: 5x4]
*/

import(
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

const (
	DayThreeID = "DayThreeID"
	Left = "l"
	Down = "d"
	WDayThreeIDth = "w"
	Height = "h"
)

func DayThreePartOne(scanner *bufio.Scanner, dimension int) int {
	square := make([][]string, dimension)
	hitMap := make(map[string]bool)

	for i := 0; i < dimension; i++ {
		square[i] = make([]string, dimension)
	}

	for scanner.Scan() {
		claimvals := parseClaim(scanner.Text())
		claimvalsint := convertClaimNumValues(claimvals)

		fillSquare(dimension, claimvals[DayThreeID], square, claimvalsint[0], claimvalsint[1], claimvalsint[2], claimvalsint[3], hitMap)
	}

	conflicts := 0
	for h := 0; h < dimension; h++ {
		for w := 0; w < dimension; w++ {
			if square[h][w] == "X" {
				conflicts++
			}
		}
	}

	// Fulfills part 2
	// for DayThreeID, _ := range hitMap {
	// 	fmt.Println("No touchie: ", DayThreeID)
	// }
	return conflicts
}

// Sample input: "#123 @ 3,2: 5x4"
func parseClaim(claim string) map[string]string {
	split := strings.Split(claim, " ")
	values := extractClaimValues(split)

	return values
}

func extractClaimValues(split []string) map[string]string {
	values := make(map [string]string)

	// Set DayThreeID
	values[DayThreeID] = split[0][1:]

	// Set left and down start values
	startPoint := strings.Split(split[2], ",")
	values[Left] = startPoint[0] 
	values[Down] = startPoint[1][:len(startPoint[1]) - 1]
	
	// Set wDayThreeIDth and height
	wDayThreeIDthHeight := strings.Split(split[3], "x")
	values[WDayThreeIDth], values[Height] = wDayThreeIDthHeight[0], wDayThreeIDthHeight[1]

	return values
}

func fillSquare(dimension int, DayThreeID string, square [][]string, left int, down int, wDayThreeIDth int, height int, hitMap map[string]bool) {
	if left < 0 || down < 0 || wDayThreeIDth <= 0 || height <= 0 {
		return
	} else if left >= dimension - 1 || down >= dimension - 1 || wDayThreeIDth > dimension || height > dimension {
		return 
	}

	hasConflict := false
	for h := 0; h < height; h++ {
		for w := 0; w < wDayThreeIDth; w++ {
			if square[left + w][down + h] != "" {
				delete(hitMap, square[left + w][down + h])
				square[left + w][down + h] = "X"
				hasConflict = true
			} else {
				square[left + w][down + h] = DayThreeID
			}
		}
	}

	if !hasConflict {
		hitMap[DayThreeID] = !hasConflict
	}
}

func convertClaimNumValues(claimvals map[string]string) [4]int {
	var intvals [4]int
	intvals[0], _ = strconv.Atoi(claimvals[Left])
	intvals[1], _ = strconv.Atoi(claimvals[Down])
	intvals[2], _ = strconv.Atoi(claimvals[WDayThreeIDth])
	intvals[3], _ = strconv.Atoi(claimvals[Height])
	return intvals
}

func printSquare(dimension int, square [][]string) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if square[i][j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(square[i][j])
			}
		}
		fmt.Println()
	}
}