package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const ResultName = "result"
const SaveName = "log"

func save(min int, count int, increase int) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Print("Cannot find program directory")
	}
	f, err := os.Create(path.Join(path.Dir(dir), SaveName))
	if err != nil {
		log.Print("Cannot create save file")
	}
	defer f.Close()
	f.Write([]byte(os.Args[1] + ";" + strconv.Itoa(min) + ";" + strconv.Itoa(count) + ";" + strconv.Itoa(increase)))
}

func readAndGetParams() (string, int, int, int) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panic("Cannot find the program directory.")
	}
	f, err := os.Open(path.Join(path.Dir(dir), SaveName))
	if err != nil {
		log.Panic("Cannot find the save file.")
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	strSaved := scanner.Text()
	splitStr := strings.Split(strSaved, ";")
	imgName := splitStr[0]
	min := splitStr[1]
	count := splitStr[2]
	increase := splitStr[3]

	resultMin, err := strconv.Atoi(min)
	if err != nil {
		resultMin = 1
	}
	resultCount, err := strconv.Atoi(count)
	if err != nil {
		resultCount = 1
	}
	resultIncrease, err := strconv.Atoi(increase)
	if err != nil {
		resultIncrease = 1
	}

	return imgName, resultMin, resultCount, resultIncrease
}
