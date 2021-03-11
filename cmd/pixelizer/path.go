package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
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
		fileDirectory = imageDir
	} else {
		fileDirectory = workingDirectory + imageDir
	}
}

func setPreviousWorkingDirectory() {
	logIfVerbose(DEBUG, "setPreviousWorkingDirectory")
	workingDirectory = os.Args[len(os.Args)-1]
}

func setWorkingDirectory() {
	logIfVerbose(DEBUG, "setWorkingDirectory")
	var errGettingPath error
	workingDirectory, errGettingPath = os.Getwd()
	logIfExists(errGettingPath)
	if errGettingPath != nil {
		return
	}
	workingDirectory += "/"
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
	}

	return split[0]

}

func setSaveDir() {
	logIfVerbose(DEBUG, "setSaveDir")
	usr, err := user.Current()
	if err != nil {
		logrus.Error(err)
	}

	saveDir = filepath.Join(usr.HomeDir, appname)

	createPixelizerDir()
}

func addWorkingDirToArgs() {
	logIfVerbose(DEBUG, "addWorkingDirToArgs")
	os.Args = append(os.Args, workingDirectory)
}

func createPixelizerDir() {
	logIfVerbose(DEBUG, "createPixelizerDir")
	_, err := os.Stat(saveDir)
	if err != nil {
		os.Mkdir(saveDir, os.ModePerm)
	}
}
