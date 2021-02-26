package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "man" || os.Args[1] == "help" {
		log.Print("Use : " + os.Args[0] + " [1] [2] [3] [4]")
		log.Print("1 : Name of the file")
		log.Print("2 : Pixelization level")
		log.Print("3 : Optional - Number of results with increasing pixelization")
		log.Print("4 : Optional - Pixelization level increase between each result")
		return
	}
	mainParam := os.Args[1]
	if mainParam == "clear" {
		clearResults()
	} else if mainParam == "redo" {
		img, min, count, increase := readAndGetParams()
		f, err := os.Open(img)
		if err != nil {
			log.Print("Cannot find this file.")
			return
		}
		launchPixellisator(f, min, count, increase)
	} else {
		f, err := os.Open(mainParam)
		if err != nil {
			log.Print("Cannot find this file.")
			return
		}
		defer f.Close()
		min := param(2)
		count := param(3)
		increase := param(4)
		launchPixellisator(f, min, count, increase)
		save(min, count, increase)
	}
}

func param(index int) int {
	var nb int
	var err error
	if len(os.Args) < (index + 1) {
		return 1
	}
	if os.Args[index] != "" {
		nb, err = strconv.Atoi(os.Args[index])
		if err != nil {
			log.Print("Cannot convert parameter " + string(index) + " to int")
		}
	} else {
		nb = 1
	}
	return nb
}
