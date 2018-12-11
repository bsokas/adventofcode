package main

import (
	"days"
	"fmt"
	"os"
	"tools"
)

const defaultPath = "../input/"

func main() {
	path := defaultPath + os.Args[1]
	fmt.Println("Running programing for file: ", path)
	if scanner, err := tools.FileScanner(path); err != nil {
		fmt.Println(err.Error())
	} else {
		answer := days.DayFive(true, scanner)
		fmt.Println("Solution: ", answer)
	}

}
