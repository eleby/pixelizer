package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
	"strconv"
)

const GifDelayEachFrame = 20

func launchPixellisator(f *os.File, min int, count int, increase int) {
	img, _, errImageDecode := image.Decode(f)
	if errImageDecode != nil {
		log.Print("Cannot decode image")
		return
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
		out, errOutput := os.Create(name)
		if errOutput != nil {
			log.Print("Cannot create image")
		}
		errOutputEncode := png.Encode(out, newImg)
		if errOutputEncode != nil {
			log.Print("Cannot encode image")
		}
		if hasParam("gif") {
			addToGif(&gifImg, newImg)
		}
		if i+1 == count {
			if hasParam("print") {
				printInTerminal(newImg, pointMin, pointMax, min+(i*increase))
			}
			if hasParam("gif") {
				gifOutput, errGif := os.Create("results.gif")
				if errGif != nil {
					log.Print("Cannot create gif image")
				}
				errGifEncode := gif.EncodeAll(gifOutput, &gifImg)
				if errGifEncode != nil {
					log.Print("Cannot encode gif image")
				}
			}
		}
	}
}

func repixellise(img image.Image, pointMin image.Point, pointMax image.Point, min int) image.Image {
	result := image.NewRGBA(image.Rectangle{Min: pointMin, Max: pointMax})
	for y := pointMin.Y; y < pointMax.Y; y += min {
		for x := pointMin.X; x < pointMax.X; x += min {
			colorRGBA := avg(img, x, y, min)
			squareFill(result, x, y, colorRGBA, min)
		}
	}
	return result
}

func avg(img image.Image, x int, y int, count int) color.RGBA {
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
	for i := 0; i < min; i++ {
		for j := 0; j < min; j++ {
			img.Set(x+i, y+j, rgba)
		}
	}
}

func addToGif(gifImg *gif.GIF, img image.Image) {
	palettedImage := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Src)
	gifImg.Image = append(gifImg.Image, palettedImage)
	gifImg.Delay = append(gifImg.Delay, GifDelayEachFrame)
}
