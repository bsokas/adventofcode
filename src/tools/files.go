package tools

import (
	"bufio"
	"os"
)

// var InputPath string = "../../input/day2.txt"

// FileScanner creates a Scanner with the provided file path
func FileScanner(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}

// ReadIntoStringArray reads the content of file into an array
func ReadIntoStringArray(scanner *bufio.Scanner) []string {
	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	return strings
}