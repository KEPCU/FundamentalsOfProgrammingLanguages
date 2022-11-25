package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
	"time"
)

func main() {
	path := "5.jpg"
	path2 := "6.jpg"

	start := time.Now()
	Sequential(path, path2)
	elapsed := time.Since(start)
	fmt.Println("Sequential time: ", elapsed)
}

func Sequential(path string, path2 string) {
	imagePath, err := os.Open(path)
	defer imagePath.Close()
	myImage, _, err := image.Decode(imagePath)

	imagePath2, err := os.Open(path2)
	defer imagePath2.Close()
	myImage2, _, err := image.Decode(imagePath2)

	bounds := myImage2.Bounds()

	size := myImage2.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := myImage2.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			r := uint8(math.Min(255, float64(originalColor.R)*3.5))
			g := uint8(math.Min(255, float64(originalColor.G)*3.5))
			b := uint8(math.Min(255, float64(originalColor.B)*3.5))

			c := color.RGBA{
				R: r, G: g, B: b, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	union := image.NewAlpha(bounds)

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			union.SetAlpha(x, y, color.Alpha{uint8(75)})
		}
	}

	myImage3 := image.NewRGBA(bounds)
	draw.Draw(myImage3, myImage3.Bounds(), myImage, image.ZP, draw.Src)
	draw.DrawMask(myImage3, bounds, wImg, image.ZP, union, image.ZP, draw.Over)

	result, _ := os.Create("resultSequential1.jpg")
	defer result.Close()
	err = jpeg.Encode(result, myImage3, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
