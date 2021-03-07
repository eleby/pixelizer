package main

import (
	"bufio"
	"os"
	"strings"
)

const ResultName = "result"
const SaveName = "log"
const TerminalPrintName = "print"

func save() {
	logIfVerbose(DEBUG, "save")
	f, errCreateFile := os.Create(SaveDir + SaveName)
	logIfExists(errCreateFile)
	defer f.Close()
	strSave := ""

	for i := range os.Args {
		strSave += os.Args[i]
		if i < len(os.Args)-1 {
			strSave += ";"
		}
	}
	f.Write([]byte(strSave))
}

func saveStr(str string, suffix string) {
	logIfVerbose(DEBUG, "saveStr")
	logIfVerbose(INFO, "Saving color codes to "+SaveDir+TerminalPrintName+suffix+".log")
	f, errCreateFile := os.Create(SaveDir + "/" + TerminalPrintName + suffix + ".log")
	logIfExists(errCreateFile)
	defer f.Close()
	f.Write([]byte(str))
}

func readAndGetParams() {
	logIfVerbose(DEBUG, "readAndGetParams")
	f, errOpenSaveDir := os.Open(SaveDir + SaveName)
	logIfExists(errOpenSaveDir)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	strSaved := scanner.Text()
	splitStr := strings.Split(strSaved, ";")

	os.Args = []string{}
	for i := range splitStr {
		os.Args = append(os.Args, splitStr[i])
	}
}
