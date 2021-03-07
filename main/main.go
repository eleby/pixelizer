package main

import (
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var WorkingDirectory string
var FileDirectory string
var SaveDir string
var AppName = "pixelizer"

func main() {
	if len(os.Args) < 2 || os.Args[1] == "man" || os.Args[1] == "help" {
		printManual()
		return
	}
	initLogLevel()
	mainParam := os.Args[1]

	if mainParam == "clear" {
		setSaveDir()
		readAndGetParams()
		setPreviousWorkingDirectory()
		addWorkingDirToArgs()
		setFileDirectory()
		clearResults()
	} else if mainParam == "redo" {
		setSaveDir()
		readAndGetParams()
		setPreviousWorkingDirectory()
		addWorkingDirToArgs()
		setFileDirectory()
		img := os.Args[1]
		min := param(2)
		count := param(3)
		increase := param(4)
		f, errOpenImg := os.Open(FileDirectory + getNameOfFile(img))
		logIfExists(errOpenImg)
		if errOpenImg != nil {
			return
		}
		launchPixellisator(f, min, count, increase)
	} else {
		setDirVariables()
		os.Args = append(os.Args, WorkingDirectory)
		f, errOpenImage := os.Open(mainParam)
		logIfExists(errOpenImage)
		if errOpenImage != nil {
			return
		}
		defer f.Close()
		min := param(2)
		count := param(3)
		increase := param(4)
		launchPixellisator(f, min, count, increase)
		save()
	}
}
