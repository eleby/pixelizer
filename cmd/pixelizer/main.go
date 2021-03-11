package main

import (
	_ "image/jpeg"
	_ "image/png"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	workingDirectory string
	fileDirectory    string
	saveDir          string
)

func main() {
	printVersion()

	if len(os.Args) < 2 || os.Args[1] == "man" || os.Args[1] == "help" {
		printManual()
		return
	}

	initLogLevel()

	mainParam := os.Args[1]

	switch mainParam {
	case "clear":
		setSaveDir()
		readAndGetParams()
		setPreviousWorkingDirectory()
		addWorkingDirToArgs()
		setFileDirectory()
		clearResults()
	case "redo":
		setSaveDir()
		readAndGetParams()
		setPreviousWorkingDirectory()
		addWorkingDirToArgs()
		setFileDirectory()

		img := os.Args[1]
		min := param(2)
		count := param(3)
		increase := param(4)

		f, err := os.Open(fileDirectory + getNameOfFile(img))
		if err != nil {
			log.Error(err)
			return
		}

		if err := launchPixellisator(f, min, count, increase); err != nil {
			log.Error(err)
			return
		}

	default:
		setDirVariables()

		os.Args = append(os.Args, workingDirectory)

		f, err := os.Open(mainParam)
		if err != nil {
			log.Error(err)

			return
		}

		defer f.Close()

		min := param(2)
		count := param(3)
		increase := param(4)

		if err := launchPixellisator(f, min, count, increase); err != nil {
			log.Error(err)
			return
		}

		save()
	}
}
