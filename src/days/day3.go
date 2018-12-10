package days

/**
	Day 3 involves finding conflicting rectangles on an area of fabric
	Input offers dimensions of [# points inward from left, # points down from top, width, height]
	Sample: [#123 @ 3,2: 5x4]
*/

import(
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

const (
	ID = "id"
	Left = "l"
	Down = "d"
	Width = "w"
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

		fillSquare(dimension, claimvals[ID], square, claimvalsint[0], claimvalsint[1], claimvalsint[2], claimvalsint[3], hitMap)
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
	// for id, _ := range hitMap {
	// 	fmt.Println("No touchie: ", id)
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

	// Set ID
	values[ID] = split[0][1:]

	// Set left and down start values
	startPoint := strings.Split(split[2], ",")
	values[Left] = startPoint[0] 
	values[Down] = startPoint[1][:len(startPoint[1]) - 1]
	
	// Set width and height
	widthHeight := strings.Split(split[3], "x")
	values[Width], values[Height] = widthHeight[0], widthHeight[1]

	return values
}

func fillSquare(dimension int, id string, square [][]string, left int, down int, width int, height int, hitMap map[string]bool) {
	if left < 0 || down < 0 || width <= 0 || height <= 0 {
		return
	} else if left >= dimension - 1 || down >= dimension - 1 || width > dimension || height > dimension {
		return 
	}

	hasConflict := false
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if square[left + w][down + h] != "" {
				delete(hitMap, square[left + w][down + h])
				square[left + w][down + h] = "X"
				hasConflict = true
			} else {
				square[left + w][down + h] = id
			}
		}
	}

	if !hasConflict {
		hitMap[id] = !hasConflict
	}
}

func convertClaimNumValues(claimvals map[string]string) [4]int {
	var intvals [4]int
	intvals[0], _ = strconv.Atoi(claimvals[Left])
	intvals[1], _ = strconv.Atoi(claimvals[Down])
	intvals[2], _ = strconv.Atoi(claimvals[Width])
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