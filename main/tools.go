package main

import (
	"fmt"
	"log"
	"os"
)

func clearResults() {
	dirRead, err := os.Open(".")
	if err != nil {
		log.Print("Cannot open directory")
	}
	dirFiles, err := dirRead.Readdir(0)
	if err != nil {
		log.Print("Cannot read directory")
	}
	for index := range dirFiles {
		fileHere := dirFiles[index]

		name := fileHere.Name()

		isSame := true
		for i := range ResultName {
			if name[i] != ResultName[i] {
				isSame = false
			}
		}
		if isSame {
			os.Remove(name)
			fmt.Println("Removed file:", name)
		}
	}
}
