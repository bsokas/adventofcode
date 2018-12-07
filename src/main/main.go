package main

import (
	"os"
	"fmt"
	"tools"
)

const defaultPath = "../input/"

func main(){
	path := defaultPath + os.Args[1]
	fmt.Println("Running programing for file: ", path)
	if scanner , err := tools.FileScanner(path); err != nil {
		fmt.Println(err.Error())
	}

}