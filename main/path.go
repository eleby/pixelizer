package main

import (
	"os"
	"os/user"
	"strings"
)

func setDirVariables() {
	logIfVerbose(DEBUG, "setDirVariables")
	setSaveDir()
	setWorkingDirectory()
	addWorkingDirToArgs()
	setFileDirectory()
}

func setFileDirectory() {
	logIfVerbose(DEBUG, "setFileDirectory")
	imageDir := getPathOf(os.Args[1])
	if imageDir[0] == '/' {
		FileDirectory = imageDir
	} else {
		FileDirectory = WorkingDirectory + imageDir
	}
}

func setPreviousWorkingDirectory() {
	logIfVerbose(DEBUG, "setPreviousWorkingDirectory")
	WorkingDirectory = os.Args[len(os.Args)-1]
}

func setWorkingDirectory() {
	logIfVerbose(DEBUG, "setWorkingDirectory")
	var errGettingPath error
	WorkingDirectory, errGettingPath = os.Getwd()
	logIfExists(errGettingPath)
	if errGettingPath != nil {
		return
	}
	WorkingDirectory += "/"
}

func getPathOf(filePath string) string {
	logIfVerbose(DEBUG, "getPathOf[ %v ]", filePath)
	split := strings.Split(filePath, "/")
	str := ""
	if len(split) > 1 {
		if filePath[0] == '/' {
			str += "/"
		}
		for i := range split {
			if i < len(split)-1 {
				str += split[i]
				str += "/"
			}
		}
	} else {
		str = "./"
	}
	return str
}

func getNameOfFile(filePath string) string {
	logIfVerbose(DEBUG, "getNameOfFile[ %v ]", filePath)
	split := strings.Split(filePath, "/")
	if len(split)-1 >= 0 {
		return split[len(split)-1]
	} else {
		return split[0]
	}
}

func setSaveDir() {
	logIfVerbose(DEBUG, "setSaveDir")
	usr, errGetUser := user.Current()
	logIfExists(errGetUser)
	SaveDir = usr.HomeDir + "/" + AppName + "/"
	createPixelizerDir()
}

func addWorkingDirToArgs() {
	logIfVerbose(DEBUG, "addWorkingDirToArgs")
	os.Args = append(os.Args, WorkingDirectory)
}

func createPixelizerDir() {
	logIfVerbose(DEBUG, "createPixelizerDir")
	_, err := os.Stat(SaveDir)
	if err != nil {
		os.Mkdir(SaveDir, os.ModePerm)
	}
}
