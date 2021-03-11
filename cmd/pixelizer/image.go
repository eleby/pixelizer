package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// GifDelayEachFrame is a delay between each gif frame.
const GifDelayEachFrame = 20

func launchPixellisator(f *os.File, min int, count int, increase int) error {
	logIfVerbose(DEBUG, "launchPixellisator[ %v / %v / %v / %v ]", f.Name(), min, count, increase)
	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("image decode: %w", err)
	}

	b := img.Bounds()
	minX := b.Min.X
	minY := b.Min.Y
	maxX := b.Max.X
	maxY := b.Max.Y
	pointMin := image.Point{X: minX, Y: minY}
	pointMax := image.Point{X: maxX, Y: maxY}
	name := "result.png"
	var newImg image.Image
	gifImg := gif.GIF{}
	for i := 0; i < count; i++ {
		newImg = repixellise(img, pointMin, pointMax, min+(i*increase))
		if count > 1 {
			name = "result" + strconv.Itoa(i+1) + ".png"
		}
		out, err := os.Create(fileDirectory + name)
		if err != nil {
			log.Error(err)
		}

		if err := png.Encode(out, newImg); err != nil {
			log.Error(err)
		}

		if hasParam("gif") && !hasParam("pixel") {
			addToGif(&gifImg, newImg)
		}

		if hasParam("print") {
			printInTerminal(newImg, min+(i*increase), i)
		}

		if i+1 == count && hasParam("gif") && !hasParam("pixel") {
			if hasParam("reverse") {
				reverseGif(&gifImg)
			}

			if hasParam("full") {
				addReverseToGif(&gifImg)
			}

			gifOutput, err := os.Create(fileDirectory + "results.gif")
			if err != nil {
				log.Error(err)
			}

			if err := gif.EncodeAll(gifOutput, &gifImg); err != nil {
				log.Error(err)
			}
		}
	}

	return nil
}

func repixellise(img image.Image, pointMin image.Point, pointMax image.Point, pixelization int) image.Image {
	logIfVerbose(DEBUG, "repixellise[ %v / %v / %v ]", pointMin, pointMax, pixelization)

	i := 0
	j := 0

	var result *image.RGBA

	if hasParam("pixel") {
		newPointMax := image.Point{X: (pointMax.X / pixelization) + 1, Y: (pointMax.Y / pixelization) + 1}
		result = image.NewRGBA(image.Rectangle{Min: pointMin, Max: newPointMax})
	} else {
		result = image.NewRGBA(image.Rectangle{Min: pointMin, Max: pointMax})
	}

	for y := pointMin.Y; y < pointMax.Y; y += pixelization {
		for x := pointMin.X; x < pointMax.X; x += pixelization {
			colorRGBA := avg(img, x, y, pixelization)
			if hasParam("pixel") {
				result.Set(i, j, colorRGBA)
			} else {
				squareFill(result, x, y, colorRGBA, pixelization)
			}
			i++
		}
		i = 0
		j++
	}

	return result
}

func avg(img image.Image, x int, y int, count int) color.RGBA {
	logIfVerbose(INFO, "avg[ { %v ; %v } / %v ]", x, y, count)

	total := count * count

	r := 0
	g := 0
	b := 0
	a := 0

	var rTmp uint32
	var gTmp uint32
	var bTmp uint32
	var aTmp uint32

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			rTmp, gTmp, bTmp, aTmp = img.At(x+i, y+j).RGBA()

			r += int(rTmp / 0x101)
			g += int(gTmp / 0x101)
			b += int(bTmp / 0x101)
			a += int(aTmp / 0x101)
		}
	}

	rAvg := r / total
	gAvg := g / total
	bAvg := b / total
	aAvg := a / total

	return color.RGBA{R: uint8(rAvg), G: uint8(gAvg), B: uint8(bAvg), A: uint8(aAvg)}
}

func squareFill(img *image.RGBA, x int, y int, rgba color.RGBA, min int) {
	logIfVerbose(INFO, "squareFill[ { %v ; %v } / %v / %v ]", x, y, rgba, min)
	for i := 0; i < min; i++ {
		for j := 0; j < min; j++ {
			img.Set(x+i, y+j, rgba)
		}
	}
}

func addToGif(gifImg *gif.GIF, img image.Image) {
	logIfVerbose(DEBUG, "addToGif")
	palettedImage := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Src)
	gifImg.Image = append(gifImg.Image, palettedImage)
	gifImg.Delay = append(gifImg.Delay, GifDelayEachFrame)
}

func addReverseToGif(gifImg *gif.GIF) {
	logIfVerbose(DEBUG, "addReverseToGif")

	for i := len(gifImg.Image) - 2; i > 0; i-- {
		gifImg.Image = append(gifImg.Image, gifImg.Image[i])
		gifImg.Delay = append(gifImg.Delay, GifDelayEachFrame)
	}
}

func reverseGif(gifImg *gif.GIF) {
	logIfVerbose(DEBUG, "reverseGif")

	var images = make([]*image.Paletted, 0, len(gifImg.Image))

	for i := len(gifImg.Image) - 1; i >= 0; i-- {
		images = append(images, gifImg.Image[i])
	}

	gifImg.Image = images
}
