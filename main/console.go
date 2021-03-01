package main

import (
	"image"
	"image/color/palette"
	"image/draw"
	"log"
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

func printInTerminal(img image.Image, pointMin image.Point, pointMax image.Point, jump int) {
	bounds := img.Bounds()
	changedImage := image.NewPaletted(bounds, palette.Plan9)
	draw.Draw(changedImage, changedImage.Rect, img, bounds.Min, draw.Src)
	printedText := "\n"
	for y := pointMin.Y; y < pointMax.Y; y += jump {
		for x := pointMin.X; x < pointMax.X; x += jump {
			colorA := changedImage.ColorIndexAt(x, y)
			printedText += "\033[48;5;" + strconv.Itoa(int(colorA)) + "m  "
		}
		printedText += "\033[0;00m\n"
	}
	log.Print(printedText)
	saveStr(printedText)
}
