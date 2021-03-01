package main

import (
	"os"
	"strconv"
)

func param(index int) int {
	var nb int
	var err error
	if len(os.Args) < (index + 1) {
		return 1
	}
	if os.Args[index] != "" {
		nb, err = strconv.Atoi(os.Args[index])
		if err != nil {
			nb = 1
		}
	} else {
		nb = 1
	}
	return nb
}

func hasParam(param string) bool {
	for i := range os.Args {
		if os.Args[i] == param {
			return true
		}
	}
	return false
}
