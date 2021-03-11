package main

import (
	"os"
)

func clearResults() {
	logIfVerbose(DEBUG, "clearResults")
	removeFilesInDir(fileDirectory, resultName)
	removeFilesInDir(saveDir, terminalPrintName)
}

func removeFilesInDir(dir string, filesPrefix string) {
	logIfVerbose(DEBUG, "removeFilesInDir[ %v / %v ]", dir, filesPrefix)
	dirRead, errOpenDir := os.Open(dir)
	logIfExists(errOpenDir)
	dirFiles, errReadDir := dirRead.Readdir(0)
	logIfExists(errReadDir)
	for index := range dirFiles {
		fileInDir := dirFiles[index]
		name := fileInDir.Name()
		isSame := true
		for i := range filesPrefix {
			if len(name) < len(filesPrefix) || name[i] != filesPrefix[i] {
				isSame = false
			}
		}
		if isSame {
			errRemoveFile := os.Remove(dir + name)
			logIfExists(errRemoveFile)
			if errRemoveFile == nil {
				logIfVerbose(DEBUG, "Removed file: %v", name)
			}
		}
	}
}
